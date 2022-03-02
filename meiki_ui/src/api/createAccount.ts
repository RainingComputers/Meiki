import { makeRequest } from "./request"

export const createAccount = (username: string, password: string) => {
    const url = "http://localhost:8080"
    const body = {
        username,
        password,
    }
    makeRequest(url, "POST", body)
}
