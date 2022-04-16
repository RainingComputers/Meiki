/// <reference types="cypress"/>

describe("Layout Test", () => {
    beforeEach(() => {
        cy.testRequest("delete", false)
        cy.clearLocalStorage()
        cy.testRequest("create", true)
        cy.testRequest("login", true)
    })

    it("App should have proper layout", () => {
        cy.visit("/notes")
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='explorer']").should("be.visible")
    })

    it("Explorer should toggle in and out", () => {
        cy.visit("/notes")

        cy.get("[data-cy='explorer']").should("be.visible")
        cy.get("[data-cy='sidebar']").should("have.class", "isChecked").click()
        cy.get("[data-cy='explorer']").should("not.exist")
        cy.get("[data-cy='sidebar']").click()
        cy.get("[data-cy='explorer']").should("be.visible")
    })

    it("Workbench show show rendered and editor accordingly", () => {
        cy.visit("/notes")

        // both should not be visible when app has been opened
        cy.get("[data-cy='editor']").should("not.exist")
        cy.get("[data-cy='renderer']").should("not.exist")

        // open both editor and renderer
        cy.get("[data-cy='edit']").should("not.have.class", "isChecked").click()
        cy.get("[data-cy='render']")
            .should("not.have.class", "isChecked")
            .click()

        // both editor and render should be visible
        cy.get("[data-cy='editor']").should("be.visible")
        cy.get("[data-cy='renderer']").should("be.visible")

        // close editor
        cy.get("[data-cy='edit']").should("have.class", "isChecked").click()

        // editor should not be visible but the renderer should be visible
        cy.get("[data-cy='editor']").should("not.exist")
        cy.get("[data-cy='renderer']").should("be.visible")

        // open editor again
        cy.get("[data-cy='edit']").should("not.have.class", "isChecked").click()

        // both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // close renderer
        cy.get("[data-cy='render']").should("have.class", "isChecked").click()

        // editor should be visible but the renderer should not be visible
        cy.get("[data-cy='renderer']").should("not.exist")
        cy.get("[data-cy='editor']").should("be.visible")

        // open renderer again
        cy.get("[data-cy='render']")
            .should("not.have.class", "isChecked")
            .click()

        // both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")
    })
})
