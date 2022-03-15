/// <reference types="cypress"/>

const baseUrl = "http://localhost:3000/"

describe("UserFlow Test", () => {


    beforeEach(() => {
        cy.testRequest("delete", false)
    })

    it("Login flow works fully", () => {
        cy.clearCookies()
        cy.visit(baseUrl)

        // shows login page
        cy.get("img[alt='meiki-logo']").should("be.visible")
        cy.get("#Username").should("be.visible")
        cy.get("#Password").should("be.visible")
        cy.get("Button").should("include.text", "Login").and("be.visible")
        cy.get("a").should("include.text", "Create").and("be.visible").click()

        // user creates an account
        cy.get("#Username").type("shnoo")
        cy.get("#Password").type("thisisveryunsafe")
        cy.get("Button")
            .should("include.text", "Create Meiki account")
            .and("be.visible")
            .click()

        // goes to account creation success page
        cy.contains("Your account has successfully been created").should(
            "be.visible"
        )
        cy.get("a").should("be.visible").click()

        // user logs in
        cy.get("#Username").type("shnoo")
        cy.get("#Password").type("thisisveryunsafe")
        cy.get("Button").click()

        // assert it goes to the app
        cy.get("nav").should("be.visible")
        // TODO: adder username in the app
    })

    it("App should have proper layout", () => {
        // create user
        cy.testRequest("create", true)

        // login
        cy.testRequest("login", true)

        cy.visit(baseUrl)

        cy.get("nav").should("be.visible")
        cy.get("[data-cy='explorer']").should("be.visible")
    })
})
