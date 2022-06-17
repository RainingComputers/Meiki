/// <reference types="cypress"/>

const DEBOUNCE_INTERVAL = 1000
const EPSILON = 100

describe("Note is saved and synchronized on changes", () => {
    beforeEach(() => {
        cy.login()
    })

    it("Note has a minimum sync interval", () => {
        const testContent = "This is a content is not synced"

        cy.visit("/")

        // Create test note and edit it, the editor will try to sync every 70 ms
        cy.createNote("testNote")
        cy.contains("testNote").click()
        cy.get(".ace_text-input")
            .first()
            .focus()
            .type(testContent, { delay: 70 })

        // assert sync indicator
        cy.get("nav")
            .get("[data-cy='badge']")
            .find("svg")
            .should("have.class", "feather-refresh-cw")

        // Simulate refreshing the page
        cy.visit("/")

        // Assert the text file was not saved since the editor was never idle for DEBOUNCE_INTERVAL ms
        // This tests weather the debouncer only calls sync only if the app was idle for a period of time
        cy.contains("testNote").click()
        cy.get("[data-cy='renderer']").should("contain.text", "")
    })

    it("Note should be saved if the user waited for enough time", () => {
        const testContent =
            "This is a content to test if the note is synced and saved properly"

        cy.visit("/")

        // Create test note and edit it
        cy.createNote("testNote")
        cy.contains("testNote").click()
        cy.get(".ace_text-input").first().focus().type(testContent)

        // assert sync indicator
        cy.get("nav")
            .get("[data-cy='badge']")
            .find("svg")
            .should("have.class", "feather-refresh-cw")

        // Wait for some time, let the app be idle
        cy.wait(DEBOUNCE_INTERVAL + EPSILON)
        cy.get("nav")
            .get("[data-cy='badge']")
            .find("svg")
            .should("have.class", "feather-check")

        // Refresh the page
        cy.visit("/")

        // Assert the note was saved
        // Since the app was idle for DEBOUNCE_INTERVAL, a sync should have happened
        cy.contains("testNote").click()
        cy.get("[data-cy='renderer']").should("contain.text", testContent)
    })

    it("Note saves on deselecting the note", () => {
        const testContent =
            "This is a content to test if the note is synced and saved properly"

        cy.visit("/")

        // Create test note and edit it
        cy.createNote("testNote")
        cy.contains("testNote").click()
        cy.get(".ace_text-input").first().focus().type(testContent)

        // Deselect the note
        cy.contains("[data-cy='explorer']", "testNote").click()

        // Refresh app
        cy.visit("/")

        // Assert the note was saved
        cy.contains("testNote").click()
        cy.get("[data-cy='renderer']").should("contain.text", testContent)
    })

    it("Errors out with unable to connect to server", () => {
        const testContent =
            "This is a content to test if sync error is handled correctly"

        cy.visit("/")
        cy.simulateServerDown("/notes/update/*")

        // Create test note and edit it
        cy.createNote("testNote")
        cy.contains("testNote").click()
        cy.get(".ace_text-input").first().focus().type(testContent)

        // Wait for some time, let the app be idle
        cy.wait(DEBOUNCE_INTERVAL + EPSILON)

        // Error should have appeared
        cy.get("nav").should("contain.text", "SYNC ERROR")
        cy.get("nav")
            .get("[data-cy='badge']")
            .find("svg")
            .should("have.class", "feather-alert-triangle")

        // Explorer should show toast
        cy.get("[data-cy='explorer']").should(
            "contain",
            "An error has occurred while syncing note, unable to connect to server"
        )
    })
})
