import { makeRequest } from "./request"

export async function login(username: string, password: string) {
    const body = {
        username,
        password,
    }

    const response = await makeRequest("/login", "POST", body)

    localStorage.setItem("username", response.username)
    localStorage.setItem("token", response.token)
}
