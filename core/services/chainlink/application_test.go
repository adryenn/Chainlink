// +build !windows

package chainlink_test

import (
	"syscall"
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/stretchr/testify/assert"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	"github.com/tevino/abool"
)

func TestChainlinkApplication_SignalShutdown(t *testing.T) {
	ethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		ethClient,
	)
	defer cleanup()
	completed := abool.New()
	app.Exiter = func(code int) {
		completed.Set()
	}

	require.NoError(t, app.Start())
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	gomega.NewGomegaWithT(t).Eventually(func() bool {
		return completed.IsSet()
	}).Should(gomega.BeTrue())
}

func TestChainlinkApplication_resumesPendingConnection_Happy(t *testing.T) {
	ethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		ethClient,
	)
	defer cleanup()
	store := app.Store

	j := cltest.NewJobWithWebInitiator()
	require.NoError(t, store.CreateJob(&j))

	jr := cltest.CreateJobRunWithStatus(t, store, j, models.RunStatusPendingConnection)

	require.NoError(t, app.StartAndConnect())
	_ = cltest.WaitForJobRunToComplete(t, store, jr)
}

func TestChainlinkApplication_resumesPendingConnection_Archived(t *testing.T) {
	ethClient, _, assertMocksCalled := cltest.NewEthMocksWithStartupAssertions(t)
	defer assertMocksCalled()
	app, cleanup := cltest.NewApplication(t,
		ethClient,
	)
	defer cleanup()
	store := app.Store

	j := cltest.NewJobWithWebInitiator()
	require.NoError(t, store.CreateJob(&j))

	jr := cltest.CreateJobRunWithStatus(t, store, j, models.RunStatusPendingConnection)

	require.NoError(t, store.ArchiveJob(j.ID))

	require.NoError(t, app.StartAndConnect())
	_ = cltest.WaitForJobRunToComplete(t, store, jr)
}

func TestChainlinkApplication_ChangeInChainID(t *testing.T) {
	config, cfgCleanup := cltest.NewConfig(t)
	defer cfgCleanup()
	app, cleanup := cltest.NewApplicationWithConfig(t, config)
	defer cleanup()

	config.Set("ETH_CHAIN_ID", 7853)
	require.NoError(t, app.Start())
	require.NoError(t, app.Stop())

	config.Set("ETH_CHAIN_ID", 2663)
	err := app.Start()
	assert.Equal(t, chainlink.ErrNewChainID, err)
}
