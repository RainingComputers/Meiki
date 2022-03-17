import { makeRequest } from "./request"

export async function login(username: string, password: string) {
    const url = "http://localhost:8080/login"

    const body = {
        username,
        password,
    }

    const response = await makeRequest(url, "POST", body)

    localStorage.setItem("username", response.username)
    localStorage.setItem("token", response.token)
}
