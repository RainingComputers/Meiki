import { makeRequest } from "./request"

export async function createAccount(username: string, password: string) {
    const body = {
        username,
        password,
    }

    await makeRequest("/create", "POST", body)
}
