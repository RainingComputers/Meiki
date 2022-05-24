/// <reference types="cypress"/>

describe("App explorer", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("App explorer should show empty notes watermark if non notes are present", () => {
        cy.visit("/")

        cy.contains(
            "Click the 'Create' button on the toolbar to create a new note"
        ).should("exist")
    })

    it("App explorer should select appropriate note", () => {
        const testContent1 =
            "This is a test note, this should be automatically saved"
        const testContent2 =
            "This is an another note, this should also be saved automatically"

        // Pre create notes
        cy.createNote("testNote1")
        cy.createNote("testNote2")
        cy.visit("/")

        // Edit note 1
        cy.contains("testNote1").click()
        cy.get("textarea").type(testContent1, { force: true })

        // Edit note 2
        cy.contains("testNote2").click()
        cy.get("textarea").type(testContent2, { force: true })

        // Select note 1 and assert contents are present
        cy.contains("testNote1").click()
        cy.get("[data-cy='renderer']").should("contain.text", testContent1)

        // Select note 2 and assert contents are present
        cy.contains("testNote2").click()
        cy.get("[data-cy='renderer']").should("contain.text", testContent2)

        // Select the panel and assert no note is selected
        cy.get("[data-cy='explorer']").click()
        cy.get("[data-cy='editor']").should("not.exist")
        cy.get("[data-cy='renderer']").should("not.exist")
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
