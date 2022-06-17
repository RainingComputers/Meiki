/// <reference types="cypress"/>

describe("Rename note", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Rename note flow", () => {
        cy.createNote("testNote")
        cy.visit("/")

        // Select the note
        cy.contains("testNote").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote").click()
        cy.get("nav").get("input").type("renamedNote").type("{enter}")

        // Assert name has changed
        cy.get("nav").contains("renamedNote")
        cy.contains("[data-cy='explorer']", "renamedNote")
    })

    it("Rename note if clicked outside label", () => {
        cy.createNote("testNote")
        cy.visit("/")

        // Select the note
        cy.contains("testNote").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote").click()
        cy.get("nav").get("input").type("renamedNote")

        // Click outside the title
        cy.get("textarea").focus()

        // Assert name has changed
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
