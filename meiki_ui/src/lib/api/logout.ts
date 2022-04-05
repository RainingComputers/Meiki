import { makeRequest } from "./request"

export async function logout() {
    await makeRequest("/logout", "POST", {})
    localStorage.clear()
}
