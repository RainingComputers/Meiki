/// <reference types="cypress"/>

describe("User account creation and login", () => {
    beforeEach(() => {
        cy.cleanUsers()
    })

    it("Login flow works fully", () => {
        cy.visit("/login")
        // Shows login page
        cy.get("[data-cy='meiki-logo']").should("be.visible")
        cy.get("#username").should("be.visible")
        cy.get("#password").should("be.visible")
        cy.get("Button").should("include.text", "Login").and("be.visible")
        cy.get("a[href='/create']").and("be.visible").click()

        // User creates an account
        cy.get("Button").should("include.text", "Create Meiki account").and("be.visible")
        cy.get("#username").type("shnoo")
        cy.get("#password").type("thisisveryunsafe")
        cy.get("#confirmPassword").type("thisisveryunsafe")
        cy.get("Button").click()

        // Goes to create success page
        cy.contains("Your account has successfully been created").should("be.visible")
        cy.get("a[href='/login']").should("be.visible").click()

        // User logs in
        cy.get("#username").type("shnoo")
        cy.get("#password").type("thisisveryunsafe")
        cy.get("Button").click()

        // Assert it goes to the app
        cy.get("nav").should("be.visible")
        cy.get("[data-cy='profile']").should("contain", "shnoo").click()

        // Click logout button
        cy.get("button:contains('Logout')").click()

        // Shows login page
        cy.get("[data-cy='meiki-logo']").should("be.visible")
        cy.get("#username").should("be.visible")
        cy.get("#password").should("be.visible")
    })

    it("Show password do not match if confirm field does not match", () => {
        cy.visit("/create")

        cy.get("#username").type("shnoo")
        cy.get("#password").type("password")
        cy.get("#confirmPassword").type("doesNotMatch")
        cy.get("Button").click()

        cy.contains("Passwords do not match").should("be.visible")
    })

    it("Show duplicate user error on create", () => {
        cy.createUser("alex1234", "password")
        cy.visit("/create")

        cy.get("#username").type("alex1234")
        cy.get("#password").type("password")
        cy.get("#confirmPassword").type("password")
        cy.get("Button").click()

        cy.contains("User already exists").should("be.visible")
    })

    it("Show invalid username error on create", () => {
        cy.visit("/create")

        cy.get("#username").type("alex**")
        cy.get("#password").type("password")
        cy.get("#confirmPassword").type("password")
        cy.get("Button").click()

        cy.contains(
            "Username should not contain any special characters or space other than '-' and '_'"
        ).should("be.visible")
    })

    it("Show invalid password error on create", () => {
        cy.visit("/create")

        cy.get("#username").type("alex1234")
        cy.get("#password").type("123")
        cy.get("#confirmPassword").type("123")
        cy.get("Button").click()

        cy.contains("Password should have minimum five characters").should("be.visible")
    })

    it("Error out with unable to connect to server on create", () => {
        cy.visit("/create")
        cy.simulateServerDown("/auth/create")

        cy.get("#username").type("alex1234")
        cy.get("#password").type("password")
        cy.get("#confirmPassword").type("password")
        cy.get("Button").click()

        cy.contains(
            "An error has occurred while creating account, unable to connect to server"
        ).should("be.visible")
    })

    it("Show invalid username error on login", () => {
        cy.visit("/login")

        cy.get("#username").type("alex**")
        cy.get("#password").type("password")
        cy.get("Button").click()

        cy.contains(
            "Username should not contain any special characters or space other than '-' and '_'"
        ).should("be.visible")
    })

    it("Show invalid password error on login", () => {
        cy.visit("/login")

        cy.get("#username").type("alex")
        cy.get("#password").type("123")
        cy.get("Button").click()

        cy.contains("Password should have minimum five characters").should("be.visible")
    })

    it("Show password mismatch on login", () => {
        cy.createUser("alex1234", "password")
        cy.visit("/login")

        cy.get("#username").type("alex1234")
        cy.get("#password").type("passwordMismatch")
        cy.get("Button").click()

        cy.contains("Password does not match").should("be.visible")
    })

    it("Error out with unable to connect to server on login", () => {
        cy.visit("/login")
        cy.simulateServerDown("/auth/login")

        cy.get("#username").type("alex1234")
        cy.get("#password").type("password")
        cy.get("Button").click()

        cy.contains("An error has occurred while logging in, unable to connect to server").should(
            "be.visible"
        )
    })
})
