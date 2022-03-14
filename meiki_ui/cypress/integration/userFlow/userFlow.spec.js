const baseUrl = "http://localhost:3000/login"

describe("UserFlow Test", () => {
    it("Goes to Create Page", () => {
        cy.visit(baseUrl)
        cy.get("img").should("have.attr", "alt").should("include", "meiki-logo")
        cy.get("input").should("have.id", "Username")
    })
})
