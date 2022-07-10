/// <reference types="cypress"/>

describe("App explorer", () => {
    beforeEach(() => {
        cy.login()
        cy.cleanNotes()
    })

    it("Show empty notes watermark if non notes are present", () => {
        cy.visit("/")

        cy.contains(
            "Click the 'Create' button on the toolbar to create a new note"
        ).should("exist")
    })

    it("Select appropriate note", () => {
        const testContent1 =
            "This is a test note, this should be automatically saved"
        const testContent2 =
            "This is an another note, this should also be saved automatically"

        // Pre create notes
        cy.createNote("testNote1")
        cy.createNote("testNote2")
        cy.visit("/")

        // Edit note 1
        cy.contains("testNote1").click()
        cy.get(".ace_text-input").first().focus().type(testContent1)

        // Edit note 2
        cy.contains("testNote2").click()
        cy.get("[data-cy='renderer']").should("not.contain.text", testContent1)
        cy.get(".ace_text-input").first().focus().type(testContent2)

        // Select note 1 and assert contents are present
        cy.get("[data-cy='explorer']").contains("testNote1").click()
        cy.get("[data-cy='renderer']").should("contain", testContent1)

        // Select note 2 and assert contents are present
        cy.get("[data-cy='explorer']").contains("testNote2").click()
        cy.get("[data-cy='renderer']").should("contain", testContent2)

        // Select the panel and assert no note is selected
        cy.get("[data-cy='explorer']").click()
        cy.get("[data-cy='editor']").should("not.exist")
        cy.get("[data-cy='renderer']").should("not.exist")
    })

    it("Error out with unable to connect to server", () => {
        cy.simulateServerDown("/notes/list")
        cy.visit("/")

        cy.get("[data-cy='explorer']").should(
            "contain.text",
            "An error has occurred while listing notes, unable to connect to server"
        )
    })

    it("Error out with read error", () => {
        cy.createNote("testNote1")
        cy.simulateServerDown("/notes/read/*")

        cy.visit("/")

        cy.get("[data-cy='explorer']").contains("testNote1").click()

        cy.get("[data-cy='explorer']").should(
            "contain",
            "An error has occurred while reading note, unable to connect to server"
        )
    })

    it("Should toggle in and out", () => {
        cy.visit("/notes")

        cy.get("[data-cy='explorer']").should("be.visible")
        cy.get("[data-cy='sidebar']").should("have.class", "isChecked").click()
        cy.get("[data-cy='explorer']").should("not.exist")
        cy.get("[data-cy='sidebar']").click()
        cy.get("[data-cy='explorer']").should("be.visible")
    })

    it("Should resize the app explorer correctly", () => {
        cy.visit("/notes")

        cy.get("[data-cy='explorer']").should("be.visible")
        cy.get("[data-cy='explorer-expander']")
            .should("be.visible")
            .trigger("mousedown")
            .trigger("mousemove", { which: 1, pageX: 500 })
        cy.get("[data-cy='explorer'").invoke("width").should("eq", 500)
        cy.get("[data-cy='explorer'").trigger("mouseup")

        cy.reload()
        cy.get("[data-cy='explorer'").invoke("width").should("eq", 500)
    })
})
