const serverUrl = "http://localhost:8080"

const testAuthCreds = {
    username: "shnoo",
    password: "thisisveryunsafe",
}

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

Cypress.Commands.add("cleanUsers", () => {
    cy.testRequest("DELETE", "/auth/delete", testAuthCreds, false)
    cy.clearLocalStorage()
})

Cypress.Commands.add("login", () => {
    cy.testRequest("DELETE", "/auth/delete", testAuthCreds, false)
    cy.clearLocalStorage()
    cy.testRequest("POST", "/auth/create", testAuthCreds, true)
    cy.testRequest("POST", "/auth/login", testAuthCreds, true)
    cy.cleanNotes()
})

Cypress.Commands.add("logout", () => {
    cy.testRequest("POST", "/auth/logout", testAuthCreds, false)
})

Cypress.Commands.add("createNote", (title) => {
    cy.testRequest("POST", "/notes/create", { title })
})

Cypress.Commands.add("createUser", (username, password) => {
    const creds = { username, password }

    cy.testRequest("DELETE", "/auth/delete", creds, false)
    cy.clearLocalStorage()
    cy.testRequest("POST", "/auth/create", creds, true)
})

Cypress.Commands.add("simulateServerDown", (endpoint) => {
    cy.intercept(endpoint, (req) => {
        req.destroy()
    })
})
