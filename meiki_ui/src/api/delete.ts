import { makeRequest } from "./request"

export async function del(username: string, password: string) {
    const url = "http://localhost:8080/delete"

    const body = {
        username,
        password,
    }

    await makeRequest(url, "POST", body)
}
