const serverUrl = "http://localhost:8080"

function getHeaders() {
    return {
        "X-Username": localStorage.getItem("username"),
        "X-Token": localStorage.getItem("token"),
    }
}

function getOptions(method, path, body = {}, failOnStatusCode = true) {
    return {
        method: method,
        url: `${serverUrl}${path}`,
        failOnStatusCode: failOnStatusCode,
        body,
        headers: getHeaders(),
    }
}

Cypress.Commands.add("testRequest", (method, path, body, failOnStatusCode) => {
    const options = getOptions(method, path, body, failOnStatusCode)

    return cy.request(options).then((response) => {
        if (path !== "/auth/login") return
        localStorage.setItem("username", response.body.username)
        localStorage.setItem("token", response.body.token)

        return response.body
    })
})

Cypress.Commands.add("cleanNotes", () => {
    const options = getOptions("GET", "/notes/list")

    return cy.request(options).then((response) => {
        response.body.forEach((noteInfo) => {
            const options = getOptions("DELETE", `/notes/delete/${noteInfo.id}`)
            cy.request(options)
        })
    })
})
