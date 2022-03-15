const serverUrl = "http://localhost:8080/"

Cypress.Commands.add("testRequest", (command, failOnStatusCode) => {
    const options = {
        method: "POST",
        url: `${serverUrl}${command}`,
        failOnStatusCode: failOnStatusCode,
        body: {
            username: "shnoo",
            password: "thisisveryunsafe",
        },
    }
    cy.request(options)
})
