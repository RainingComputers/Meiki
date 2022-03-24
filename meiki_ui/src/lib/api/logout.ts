import { ensureStatusOK, makeRequest } from "./request"

export async function logout() {
    const url = "http://localhost:8080/logout"

    await makeRequest(url, "POST", {})
    localStorage.clear()
}
