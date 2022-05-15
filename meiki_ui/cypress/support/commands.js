const serverUrl = "http://localhost:8080"

Cypress.Commands.add(
    "testRequest",
    (method, command, body, failOnStatusCode) => {
        const options = {
            method: method,
            url: `${serverUrl}${command}`,
            failOnStatusCode: failOnStatusCode,
            body,
            headers: {
                "X-Username": localStorage.getItem("username"),
                "X-Token": localStorage.getItem("token"),
            },
        }

        cy.request(options).then((response) => {
            if (command !== "/auth/login") return
            localStorage.setItem("username", response.body.username)
            localStorage.setItem("token", response.body.token)
        })
    }
)
