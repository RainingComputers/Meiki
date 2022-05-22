describe("Delete note flow", () => {
    it("Delete note flow", () => {
        /*TODO*/
        // create test note
        // select the test note
        // click the delete button
        // click the yes button
        // assert note is not present in app explorer
    })

    it("Should not create note if NO button is clicked", () => {
        /*TODO*/
        // create test note
        // select the test note
        // click the delete button
        // click the no button
        // assert note is present in app explorer
    })

    it("Should not create note if clicked outside modal", () => {
        /*TODO*/
        // create test note
        // select the test note
        // click the delete button
        // click outside the modal
        // assert note is present in app explorer
    })

    it("Should error out with unable to connect to server", () => {
        /*TODO*/
        // use cy.intercept to simulate failure on /notes/delete endpoint
        // create test note
        // select the test note
        // click the delete button
        // click the yes button
        // assert error
        // cleanup cy.intercept
    })
})
