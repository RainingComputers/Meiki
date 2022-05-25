/// <reference types="cypress"/>

describe("Hero page", () => {
    beforeEach(() => {
        cy.cleanUsers()
    })

    it("Hero page goes to create page and login page", () => {
        // Shows hero page when not logged in
        cy.visit("/")
        cy.get("span:contains('lightning fast')")

        // Goes to create page on clicking create button
        cy.get("button").click()
        cy.get("Button").should("include.text", "Create Meiki account")

        // Goes to login page on clicking login link
        cy.visit("/")
        cy.get("a:contains('Click here')").click()
        cy.get("Button").should("include.text", "Login")
    })

    it("Hero page redirects to app if logged in", () => {
        cy.login()
        cy.visit("/")

        // Goes to app if already logged in
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='profile']").should("contain", "shnoo").click()
    })
})
