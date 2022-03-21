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
        cy.get(".feather-sidebar").click()
        cy.get("[data-cy='explorer']").should("not.exist")
        cy.get(".feather-sidebar").click()
        cy.get("[data-cy='explorer']").should("be.visible")
    })
})
