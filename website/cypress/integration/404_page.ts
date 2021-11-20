describe('404 page', () => {
  it('should give the option to return to home', () => {
    cy.visit('/does-not-exists');
    cy.contains('Return to Home');
    cy.get('a').click();
    cy.contains('Learn React', { timeout: 10000 });
  });
});
