import { makeRequest } from "./request"

export async function login(username: string, password: string) {
    const url = "http://localhost:8080/login"

    const body = {
        username,
        password,
    }

    await makeRequest(url, "POST", body)
}
