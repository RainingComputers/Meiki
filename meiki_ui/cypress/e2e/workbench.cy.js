/// <reference types="cypress"/>

describe("Workbench", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("App should have proper layout", () => {
        cy.visit("/notes")
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='explorer']").should("be.visible")
    })

    it("Workbench shows rendered and editor accordingly", () => {
        cy.visit("/notes")

        // Both should not be visible when app has been opened
        cy.get("[data-cy='editor']").should("not.exist")
        cy.get("[data-cy='renderer']").should("not.exist")

        // Create note
        cy.createNote("testNote")

        // Select a note
        cy.contains("testNote").click()

        // Both editor and render should be visible
        cy.get("[data-cy='editor']").should("be.visible")
        cy.get("[data-cy='renderer']").should("be.visible")

        // Close editor
        cy.get("[data-cy='edit']").should("have.class", "isChecked").click()

        // Editor should not be visible but the renderer should be visible
        cy.get("[data-cy='editor']").should("not.be.visible")
        cy.get("[data-cy='renderer']").should("be.visible")

        // Open editor again
        cy.get("[data-cy='edit']").should("not.have.class", "isChecked").click()

        // Both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // Close renderer
        cy.get("[data-cy='render']").should("have.class", "isChecked").click()

        // Editor should be visible but the renderer should not be visible
        cy.get("[data-cy='renderer']").should("not.be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // Open renderer again
        cy.get("[data-cy='render']")
            .should("not.have.class", "isChecked")
            .click()

        // Both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")
    })
})
