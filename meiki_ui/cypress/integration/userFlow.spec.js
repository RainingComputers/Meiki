/// <reference types="cypress"/>

import { testAuthCreds } from "./testAuthCreds"

describe("UserFlow Test", () => {
    before(() => {
        cy.testRequest("DELETE", "/auth/delete", testAuthCreds, false)
        cy.clearLocalStorage()
    })

    it("Login flow works fully", () => {
        cy.visit("/login")
        // shows login page
        cy.get("img[alt='meiki-logo']").should("be.visible")
        cy.get("#username").should("be.visible")
        cy.get("#password").should("be.visible")
        cy.get("Button").should("include.text", "Login").and("be.visible")
        cy.get("a[href='/create']").and("be.visible").click()

        // user creates an account
        cy.get("Button")
            .should("include.text", "Create Meiki account")
            .and("be.visible")
        cy.get("#username").type("shnoo")
        cy.get("#password").type("thisisveryunsafe")
        cy.get("#confirmpassword").type("thisisveryunsafe")
        cy.get("Button").click()

        // goes to account creation success page
        cy.contains("Your account has successfully been created").should(
            "be.visible"
        )
        cy.get("a[href='/login']").should("be.visible").click()

        // user logs in
        cy.get("#username").type("shnoo")
        cy.get("#password").type("thisisveryunsafe")
        cy.get("Button").click()

        // assert it goes to the app
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='profile']").should("contain", "shnoo").click()

        // click logout button
        cy.get("button:contains('Logout')").click()

        // shows login page
        cy.get("img[alt='meiki-logo']").should("be.visible")
        cy.get("#username").should("be.visible")
        cy.get("#password").should("be.visible")
    })

    it("Should show password do not match if confirm field does not match", () => {
        /*TODO*/
    })

    it("Should show duplicate user error on account creation", () => {
        /*TODO*/
    })

    it("Should show invalid username error on account creation", () => {
        /*TODO*/
    })

    it("Should show invalid password error on account creation", () => {
        /*TODO*/
    })

    it("Should error out with unable to connect to server on account creation", () => {
        /*TODO*/
        // use cy.intercept to simulate failure on /auth/create endpoint
        // assert error
        // cleanup cy.intercept
    })

    it("Should show invalid username error on login", () => {
        /*TODO*/
    })

    it("Should show invalid password error on login", () => {
        /*TODO*/
    })

    it("Should show password mismatch on login", () => {
        /*TODO*/
    })

    it("Should error out with unable to connect to server on login", () => {
        /*TODO*/
        // use cy.intercept to simulate failure on /auth/login endpoint
        // assert error
        // cleanup cy.intercept
    })
})
