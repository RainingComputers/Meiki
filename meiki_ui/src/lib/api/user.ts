import { makeRequest } from "./request"

export async function createAccount(username: string, password: string) {
    const body = {
        username,
        password,
    }

    await makeRequest("/create", "POST", body)
}


export async function login(username: string, password: string) {
    const body = {
        username,
        password,
    }

    const response = await makeRequest("/login", "POST", body)

    localStorage.setItem("username", response.username)
    localStorage.setItem("token", response.token)
}

export async function authStatus() {
    await makeRequest("/authStatus", "GET")
}

export async function del(username: string, password: string) {
    const body = {
        username,
        password,
    }

    await makeRequest("/delete", "POST", body)
}

export async function logout() {
    await makeRequest("/logout", "POST")
    localStorage.clear()
}
