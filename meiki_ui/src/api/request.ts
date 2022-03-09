// from https://github.com/RainingComputers/shnootalk-playground/blob/main/ui/src/api/request.ts

function ensureStatusOK(response: Response) {
    if (response.status !== 200)
        throw { error: "response not OK", context: response }
}

export async function makeRequest(url: string, method: string, body: any = {}) {
    let requestOptions: any = { method }

    if (body)
        requestOptions = {
            method,
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
            credentials: "include",
        }

    const response = await fetch(url, requestOptions)
    ensureStatusOK(response)
    return response.json()
}
