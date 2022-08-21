/// <reference types="cypress"/>

describe("Rename note", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Rename note flow", () => {
        cy.createNote("testNote")
        cy.visit("/")

        // Select the note
        cy.contains("testNote").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote").click()
        cy.get("nav").find("input").type("renamedNote").type("{enter}")

        // Assert name has changed
        cy.get("nav").contains("renamedNote")
        cy.contains("[data-cy='explorer']", "renamedNote")
    })

    it("Rename note after another is selected should rename correctly", () => {
        cy.createNote("testNote1")
        cy.createNote("testNote2")
        cy.visit("/")

        // Select the note
        cy.contains("testNote1").click()
        cy.contains("testNote2").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote2").click()
        cy.get("nav").find("input").should("have.value", "testNote2")
    })

    it("Rename note if clicked outside label", () => {
        cy.createNote("testNote")
        cy.visit("/")

        // Select the note
        cy.contains("testNote").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote").click()
        cy.get("nav").find("input").type("renamedNote")

        // Click outside the title
        cy.get("textarea").focus()

        // Assert name has changed
        cy.get("nav").contains("renamedNote")
        cy.contains("[data-cy='explorer']", "renamedNote")
    })

    it("Error out with unable to connect to server", () => {
        cy.createNote("testNote")
        cy.visit("/")
        cy.simulateServerDown("/notes/rename/*")

        // Select the note
        cy.contains("testNote").click()

        // Click on the toolbar title and rename the note
        cy.get("nav").contains("testNote").click()
        cy.get("nav").find("input").type("renamedNote").type("{enter}")

        // Assert error
        cy.get("[data-cy='explorer']").should(
            "contain",
            "An error has occurred while renaming note, unable to connect to server"
        )
    })
})
