/// <reference types="cypress"/>

const DEBOUNCE_INTERVAL = 5000
const EPSILON = 10

describe("Note is saved and synchronized on changes", () => {
    beforeEach(() => {
        cy.login()
    })

    // it("Note auto syncs", () => {
    //     const testContent =
    //         "This is a content to test if the note is synced and saved properly"

    //     const keyboardSpeed = 100
    //     const expectedNumChars = DEBOUNCE_INTERVAL / keyboardSpeed - EPSILON

    //     cy.visit("/")

    //     cy.createNote("testNote")
    //     cy.contains("testNote").click()
    //     cy.get(".ace_text-input")
    //         .first()
    //         .focus()
    //         .type(testContent, { delay: 100 })

    //     cy.visit("/login")
    //     cy.visit("/")

    //     cy.contains("testNote").click()
    //     cy.get("[data-cy='renderer']").should(
    //         "contain.text",
    //         testContent.slice(0, expectedNumChars - 1)
    //     )
    // })

    // it("Note saves on selected outside the note", () => {
    //     const testContent =
    //         "This is a content to test if the note is synced and saved properly"

    //     cy.visit("/")

    //     cy.createNote("testNote")
    //     cy.contains("testNote").click()
    //     cy.get(".ace_text-input").first().focus().type(testContent)

    //     cy.contains("[data-cy='explorer']", "testNote").click()
    //     cy.visit("/")

    //     cy.contains("testNote").click()
    //     cy.get("[data-cy='renderer']").should("contain.text", testContent)
    // })

    it("Note should be saved if the user waited for enough time", () => {
        const testContent =
            "This is a content to test if the note is synced and saved properly"

        cy.visit("/")

        cy.createNote("testNote")
        cy.contains("testNote").click()
        cy.get(".ace_text-input")
            .first()
            .focus()
            .type(testContent, { delay: 75 })
        cy.wait(DEBOUNCE_INTERVAL + EPSILON)

        cy.visit("/login")
        cy.visit("/")

        cy.contains("testNote").click()
        cy.get("[data-cy='renderer']").should("contain.text", testContent)
    })

    it("Errors out with unable to connect to server", () => {
        // TODO
    })
})
