/// <reference types="cypress"/>

describe("Delete note", () => {
    beforeEach(() => {
        cy.on("uncaught:exception", (_, __) => false)
        cy.login()
        cy.cleanNotes()
    })

    it("Delete selected note", () => {
        cy.visit("/")
        cy.createNote("testNote")

        cy.get("[data-cy='explorer']").contains("testNote").click()
        cy.get("[data-cy='explorer']").contains("testNote").trigger("mouseenter")
        cy.get("[data-cy='delete']").click()
        cy.contains("Yes").click()
        cy.get("[data-cy='explorer']").contains("testNote").should("not.exist")
        cy.get("[data-cy='noteTitle']").should("not.exist")
    })

    it("Delete note when different selected note", () => {
        cy.visit("/")
        cy.createNote("testNote1")
        cy.createNote("testNote2")

        cy.get("[data-cy='explorer']").contains("testNote1").click()
        cy.get("[data-cy='explorer']").contains("testNote1").trigger("mouseleave")

        cy.get("[data-cy='explorer']").contains("testNote2").trigger("mouseenter")
        cy.get("[data-cy='delete']:visible").click()

        cy.contains("Yes").click()

        cy.get("[data-cy='explorer']").contains("testNote2").should("not.exist")
        cy.get("[data-cy='explorer']").contains("testNote1").should("exist")

        cy.get("[data-cy='noteTitle']").contains("testNote1")
    })

    it("Delete note when no selected note", () => {
        cy.visit("/")
        cy.createNote("testNote1")
        cy.createNote("testNote2")

        cy.get("[data-cy='explorer']").contains("testNote2").trigger("mouseenter")
        cy.get("[data-cy='delete']:visible").click()

        cy.contains("Yes").click()

        cy.get("[data-cy='explorer']").contains("testNote2").should("not.exist")
        cy.get("[data-cy='explorer']").contains("testNote1").should("exist")

        cy.get("[data-cy='noteTitle']").should("not.exist")
    })

    it("Do not delete note if NO button is clicked", () => {
        cy.visit("/")
        cy.createNote("testNote")

        cy.get("[data-cy='explorer']").contains("testNote").click()
        cy.get("[data-cy='explorer']").contains("testNote").trigger("mouseenter")
        cy.get("[data-cy='delete']").click()
        cy.contains("No").click()
        cy.get("[data-cy='explorer']").contains("testNote").should("exist")
    })

    it("Do not delete note if clicked outside modal", () => {
        cy.visit("/")
        cy.createNote("testNote")

        cy.get("[data-cy='explorer']").contains("testNote").click()
        cy.get("[data-cy='explorer']").contains("testNote").trigger("mouseenter")
        cy.get("[data-cy='delete']").click()
        cy.get("[data-cy='modalOverlay']").click()
        cy.contains("No").click()
        cy.get("[data-cy='explorer']").contains("testNote").should("exist")
    })

    it("Error out with unable to connect to server", () => {
        cy.visit("/")
        cy.simulateServerDown("/notes/delete/*")
        cy.createNote("testNote")

        cy.get("[data-cy='explorer']").contains("testNote").click()
        cy.get("[data-cy='explorer']").contains("testNote").trigger("mouseenter")
        cy.get("[data-cy='delete']").click()
        cy.contains("Yes").click()
        cy.contains(
            "An error has occurred while deleting note, unable to connect to server"
        ).should("exist")
    })
})
