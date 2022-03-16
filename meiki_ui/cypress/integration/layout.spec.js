/// <reference types="cypress"/>

describe("Layout Test", () => {
    beforeEach(() => {
        cy.testRequest("delete", false)
        cy.testRequest("create", true)
        cy.testRequest("login", true)
    })

    it("App should have proper layout", () => {
        cy.visit("/")
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='explorer']").should("be.visible")
    })
})
