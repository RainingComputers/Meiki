import { makeRequest } from "./request"

export async function createAccount(username: string, password: string) {
    const url = "http://localhost:8080/create"

    const body = {
        username,
        password,
    }

    await makeRequest(url, "POST", body)
}
