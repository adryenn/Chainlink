package chainlink

import (
	"reflect"
	"sync"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/service"
	"go.uber.org/multierr"
)

// Usage
//
// services := NewServices()
// services.AddService(promReporter)
// services.AddService(fluxmonitorSvc, RunConcurrently())
//
// services.Run()
//
// services.Shutdown(merr)

// ServiceRunner defines how a service is to be run.
type ServiceRunner struct {
	svc service.Service

	runConcurrently bool
	chStop          chan struct{}
}

type Services struct {
	// wgDone waits until all concurrent services have been closed
	wgDone sync.WaitGroup

	// Public for now until we can move health checking into here.
	Runners []ServiceRunner
}

func NewServices() *Services {
	return &Services{
		Runners: []ServiceRunner{},
	}
}

type serviceRunnerOption func(runner *ServiceRunner)

func (s *Services) Register(svc service.Service, opts ...serviceRunnerOption) {
	runner := ServiceRunner{
		svc: svc,
	}

	for _, opt := range opts {
		opt(&runner)
	}

	s.Runners = append(s.Runners, runner)
}

// RunConcurrently runs the service in a goroutine
func RunConcurrently() serviceRunnerOption {
	return func(runner *ServiceRunner) {
		runner.runConcurrently = true
		runner.chStop = make(chan struct{})
	}
}

// Run starts all services that have been registered.
func (s *Services) Run() error {
	var err error
	for _, runner := range s.Runners {
		logger.Debugw("Starting service", "serviceType", reflect.TypeOf(runner.svc))

		if runner.runConcurrently {
			s.wgDone.Add(1)
			go func() {
				defer s.wgDone.Done()

				err = runner.svc.Start()

				<-runner.chStop
			}()
		} else {
			return runner.svc.Start()
		}
	}

	// TODO - Do we need to clean up the goroutines that have already been
	// started successfully? We do not currently do this.
	if err != nil {
		return err
	}

	return nil
}

// Shutdown stops the running services in the reverse order from which they were
// started.
//
// It currently takes err as an argument becuase we want to append to the
// multierrors. When all services get moved into here, we can remove this.
func (s *Services) Shutdown(merr error) error {
	for i := len(s.Runners) - 1; i >= 0; i-- {
		runner := s.Runners[i]
		logger.Debugw("Closing service...", "serviceType", reflect.TypeOf(runner.svc))

		err := runner.svc.Close()
		merr = multierr.Append(merr, err)

		if runner.runConcurrently {
			close(runner.chStop)
		}
	}

	// Wait for all concurrent services running in go routines to finish
	// stopping.
	s.wgDone.Wait()

	return merr
}
