// from https://github.com/RainingComputers/shnootalk-playground/blob/main/ui/src/api/request.ts

export function ensureStatusOK(response: Response) {
    if (response.status !== 200){
        throw { error: "response not OK", context: response }
    }
}

export async function makeRequest(url: string, method: string, body: any = {}) {
    const username = localStorage.getItem("username")
    const token = localStorage.getItem("token")

    let requestOptions: any = {
        method,
        headers: {
            "Content-Type": "application/json",
            "X-Username": username,
            "X-Token": token,
        },
        credentials: "include",
    }

    if (method === "POST"){
        requestOptions.body = JSON.stringify(body)
    }

    const response = await fetch(url, requestOptions)
    ensureStatusOK(response)
    return response.json()
}
