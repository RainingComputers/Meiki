/// <reference types="cypress"/>

import { testAuthCreds } from "./testAuthCreds"

describe("Layout Test", () => {
    beforeEach(() => {
        cy.testRequest("DELETE", "/auth/delete", testAuthCreds, false)
        cy.clearLocalStorage()
        cy.testRequest("POST", "/auth/create", testAuthCreds, true)
        cy.testRequest("POST", "/auth/login", testAuthCreds, true)
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

        // create note
        // separate test for create workflow notes, testing only workbench logic
        cy.testRequest("POST", "/notes/create", { title: "testing" })

        // select a note
        cy.contains("testing").click()

        // both editor and render should be visible
        cy.get("[data-cy='editor']").should("be.visible")
        cy.get("[data-cy='renderer']").should("be.visible")

        // close editor
        cy.get("[data-cy='edit']").should("have.class", "isChecked").click()

        // editor should not be visible but the renderer should be visible
        cy.get("[data-cy='editor']").should("not.be.visible")
        cy.get("[data-cy='renderer']").should("be.visible")

        // open editor again
        cy.get("[data-cy='edit']").should("not.have.class", "isChecked").click()

        // both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // close renderer
        cy.get("[data-cy='render']").should("have.class", "isChecked").click()

        // editor should be visible but the renderer should not be visible
        cy.get("[data-cy='renderer']").should("not.be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // open renderer again
        cy.get("[data-cy='render']")
            .should("not.have.class", "isChecked")
            .click()

        // both editor and renderer should be visible
        cy.get("[data-cy='renderer']").should("be.visible")
        cy.get("[data-cy='editor']").should("be.visible")

        // TODO: figure out a way to clean up all notes, add a testing endpoint only available when DEBUG=true env is set?
    })
})
