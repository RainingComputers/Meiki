import { makeRequest } from "./request"

export async function createUser(username: string, password: string) {
    const body = { username, password }
    await makeRequest("auth/create", "POST", body)
}

export async function login(username: string, password: string) {
    const body = { username, password }
    const response = await makeRequest("auth/login", "POST", body)

    localStorage.setItem("username", response.username)
    localStorage.setItem("token", response.token)
}

export async function authStatus() {
    await makeRequest("auth/authStatus", "GET")
}

export async function deleteUser(username: string, password: string) {
    const body = { username, password }
    await makeRequest("auth/delete", "POST", body)
}

export async function logout() {
    await makeRequest("auth/logout", "POST")
    localStorage.removeItem("username")
    localStorage.removeItem("token")
}

export function getUsername() {
    return localStorage.getItem("username")
}

export function tokensPresent() {
    const username = localStorage.getItem("username")
    const token = localStorage.getItem("token")

    return !!(username && token)
}
