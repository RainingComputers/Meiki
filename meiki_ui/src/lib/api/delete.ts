import { makeRequest } from "./request"

export async function del(username: string, password: string) {
    const body = {
        username,
        password,
    }

    await makeRequest("/delete", "POST", body)
}
