import { makeRequest } from "./request"

export type NoteInfo = {
    id: string
    title: string
    username: string
}

export async function listNotes(): Promise<Array<NoteInfo>> {
    return await makeRequest("/notes/list", "GET")
}

export async function createNote(title: string) {
    await makeRequest("/notes/create", "POST", {title})
}
