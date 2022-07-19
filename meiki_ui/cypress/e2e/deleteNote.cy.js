/// <reference types="cypress"/>

describe("Delete note", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Delete note flow", () => {
        cy.on("uncaught:exception", (_, __) => false)

        cy.visit("/")
        cy.createNote("testNote")

        cy.contains("testNote").click()
        cy.get("[data-cy='delete']").click()
        cy.contains("Yes").click()
        cy.contains("testNote").should("not.exist")
    })

    it("Do not delete note if NO button is clicked", () => {
        cy.visit("/")
        cy.createNote("testNote")

        cy.contains("testNote").click()
        cy.get("[data-cy='delete']").click()
        cy.contains("No").click()
        cy.contains("testNote").should("exist")
    })

    it("Do not delete note if clicked outside modal", () => {
        cy.visit("/")
        cy.createNote("testNote")

        cy.contains("testNote").click()
        cy.get("[data-cy='delete']").click()
        cy.get("[data-cy='modalOverlay']").click()
        cy.contains("No").click()
        cy.contains("testNote").should("exist")
    })

    it("Error out with unable to connect to server", () => {
        cy.visit("/")
        cy.simulateServerDown("/notes/delete/*")
        cy.createNote("testNote")

        cy.contains("testNote").click()
        cy.get("[data-cy='delete']").click()
        cy.contains("Yes").click()
        cy.contains(
            "An error has occurred while deleting note, unable to connect to server"
        ).should("exist")
    })
})
