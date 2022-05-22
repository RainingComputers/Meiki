describe("Create note flow", () => {
    it("Create note flow", () => {
        /*TODO*/
        // click create button
        // type note name
        // click create button
        // assert note is present in app explorer
    })

    it("Should not create note if clicked outside modal", () => {
        /*TODO*/
        // click create button
        // type note name
        // click outside the modal
        // assert note is not present in app explorer
    })

    it("Should error out with unable to connect to server", () => {
        /*TODO*/
        // use cy.intercept to simulate failure on /notes/create endpoint
        // click create button
        // type note name
        // click create button
        // assert error
        // cleanup cy.intercept
    })
})
