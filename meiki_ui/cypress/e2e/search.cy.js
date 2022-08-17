/// <reference types="cypress"/>

describe("Search across notes", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
        cy.createNote("Note1", "This is Note 1 content")
        cy.createNote("Note2", "This is Note 2 content")
        cy.createNote("Note3", "This is Note 3 content but has Note1")
    })

    it("Search works correctly", () => {
        cy.visit("/")

        cy.get("[data-cy='searchbar']").type("Note1", { delay: 10 })

        cy.get("[data-cy='explorer']").contains("Note1").should("exist")
        cy.get("[data-cy='explorer']").contains("Note3").should("exist")
        cy.get("[data-cy='explorer']").contains("Note2").should("not.exist")
    })
})
