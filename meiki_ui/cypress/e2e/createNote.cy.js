/// <reference types="cypress"/>

describe("Rename note", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Create note flow", () => {
        cy.visit("/")

        cy.contains("Create").click()
        cy.get("#enterNoteName").type("testNote")
        cy.contains("Create note").click()
        cy.contains("testNote").should("exist")
    })

    it("Do not create note if clicked outside modal", () => {
        cy.visit("/")

        cy.contains("Create").click()
        cy.get("#enterNoteName").type("testNote")
        cy.get("[data-cy='modalOverlay']").click()
        cy.contains("testNote").should("not.exist")
    })

    it("Error out with unable to connect to server", () => {
        cy.visit("/")
        cy.simulateServerDown("/notes/create")

        cy.contains("Create").click()
        cy.get("#enterNoteName").type("testNote")
        cy.contains("Create note").click()

        cy.contains(
            "An error has occurred while creating note, unable to connect to server"
        ).should("exist")
    })
})
