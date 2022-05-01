import { makeRequest } from "./request"

const AUTH_ROUTE_PREFIX = "auth"

export async function createUser(username: string, password: string) {
    const body = { username, password }
    await makeRequest(`${AUTH_ROUTE_PREFIX}/create`, "POST", body)
}

export async function login(username: string, password: string) {
    const body = { username, password }
    const response = await makeRequest(
        `${AUTH_ROUTE_PREFIX}/login`,
        "POST",
        body
    )

    localStorage.setItem("username", response.username)
    localStorage.setItem("token", response.token)
}

export async function authStatus() {
    await makeRequest(`${AUTH_ROUTE_PREFIX}/authStatus`, "GET")
}

export async function deleteUser(username: string, password: string) {
    const body = { username, password }
    await makeRequest(`${AUTH_ROUTE_PREFIX}/delete`, "POST", body)
}

export async function logout() {
    await makeRequest(`${AUTH_ROUTE_PREFIX}/logout`, "POST")
    localStorage.clear()
}

export function getUsername() {
    return localStorage.getItem("username")
}

export function tokensPresent() {
    const username = localStorage.getItem("username")
    const token = localStorage.getItem("token")

    return !!(username && token)
}
