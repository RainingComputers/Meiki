import { makeRequest } from "./request"

export async function authStatus() {
    const url = "http://localhost:8080/authStatus"
    await makeRequest(url, "GET")
}
