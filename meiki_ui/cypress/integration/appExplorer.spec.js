describe("App explorer", () => {
    beforeEach(() => {
        cy.login()
    })

    it("App explorer should show empty notes watermark if non notes are present", () => {
        /* TODO */
    })

    it("App explorer should select appropriate note", () => {
        /* TODO */
        // create two notes
        // select note1
        // edit note1 and add contents
        // hit save button ???
        // select note2
        // edit note2 and add contents
        // select note1
        // assert note1 contents are present
        // select note2
        // assert note2 contents are present
        // select the panel
        // assert editor and renderer are not visible
    })

    it("Should error out with unable to connect to server", () => {
        /*TODO*/
        // use cy.intercept to simulate failure on /notes/list endpoint
        // assert error
        // cleanup cy.intercept
    })

    it("App explorer should toggle in and out", () => {
        cy.visit("/notes")

        cy.get("[data-cy='explorer']").should("be.visible")
        cy.get("[data-cy='sidebar']").should("have.class", "isChecked").click()
        cy.get("[data-cy='explorer']").should("not.exist")
        cy.get("[data-cy='sidebar']").click()
        cy.get("[data-cy='explorer']").should("be.visible")
    })
})
