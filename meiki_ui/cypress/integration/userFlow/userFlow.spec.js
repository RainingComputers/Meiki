/// <reference types="cypress"/>

const baseUrl = "http://localhost:3000/"

describe("UserFlow Test", () => {
    it("Shows login page if not logged in", () => {
        cy.clearCookies()
        cy.visit(baseUrl)
        
        cy.get("img[alt='meiki-logo']").should("be.visible")
        cy.get("#Username").should("be.visible")
        cy.get("#Password").should("be.visible")
    })
})
