context('End to end', function () {
  it('Deletes a job', () => {
    cy.login()

    // Create Job
    cy.clickLink('New Job')
    cy.contains('span', 'New Job').should('exist')
    cy.getJobJson().then((jobJson) => {
      cy.get('textarea[id=jobSpec]').paste(jobJson)
    })
    cy.clickButton('Create Job')
    cy.contains('p', 'Successfully created job')
      .children('a')
      .invoke('text')
      .as('jobId')

    // Delete Job
    cy.get('#created-job').click()
    cy.contains('h6', 'job spec detail').should('exist')
    cy.clickButton('Delete')
    cy.contains('h5', 'Warning').should('exist')
    cy.get('@jobId').then((jobId) => {
      cy.contains('button', `Delete ${jobId}`).click()
    })
    cy.contains('p', 'Successfully deleted job').should('exist')
  })
})
