/// <reference types="cypress"/>

describe("Create note", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Rename note flow", () => {
        cy.createNote("testNote")

        // Go to root
        cy.visit("/")

        // Click the note name
        cy.contains("testNote").click()

        // Click the note label
        cy.get("nav")
            .contains("testNote")
            .click()
            .type("{selectAll}renamedNote")
            .type("{enter}")

        // Verify that name is changed
        cy.get("nav").contains("renamedNote")
        cy.contains("[data-cy='explorer']", "renamedNote")
    })

    it("Rename note if clicked outside label", () => {
        cy.createNote("testNote")

        // Go to root
        cy.visit("/")

        // Click the note name
        cy.contains("testNote").click()

        // Click the note label
        cy.get("nav")
            .contains("testNote")
            .click()
            .type("{selectAll}renamedNote")

        cy.get("textarea").focus()

        // Verify that name is changed
        cy.get("nav").contains("renamedNote")
        cy.contains("[data-cy='explorer']", "renamedNote")
    })

    // TODO: Error
    // it("Error out in rename if unable to connect to server", () => {
    //     cy.visit("/")
    //     cy.simulateServerDown("/notes/create")

    //     cy.contains("Create").click()
    //     cy.get("#enterNoteName").type("testNote")
    //     cy.contains("Create note").click()

    //     cy.contains(
    //         "An error has occurred while creating note, unable to connect to server"
    //     ).should("exist")
    // })
})
