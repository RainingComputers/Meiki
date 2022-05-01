const serverUrl = "http://localhost:8080"

Cypress.Commands.add("testRequest", (method, command, failOnStatusCode) => {
    const options = {
        method: method,
        url: `${serverUrl}${command}`,
        failOnStatusCode: failOnStatusCode,
        body: {
            username: "shnoo",
            password: "thisisveryunsafe",
        },
    }

    const response = cy.request(options).then((response) => {
        if (command !== "/auth/login") return
        localStorage.setItem("username", response.body.username)
        localStorage.setItem("token", response.body.token)
    })
})
