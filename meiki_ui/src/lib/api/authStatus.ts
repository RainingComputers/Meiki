import { makeRequest } from "./request"

export async function authStatus() {
    await makeRequest("/authStatus", "GET")
}
