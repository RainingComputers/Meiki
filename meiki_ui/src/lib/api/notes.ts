import { makeRequest } from "./request"

export type NoteInfo = {
    id: string
    title: string
    username: string
}

export async function listNotes(): Promise<Array<NoteInfo>> {
    return await makeRequest("/notes/list", "GET")
}

export async function createNote(title: string): Promise<NoteInfo> {
    return await makeRequest("/notes/create", "POST", { title })
}

export async function deleteNote(id: string) {
    await makeRequest(`/notes/delete/${id}`, "DELETE")
}

export async function updateNote(id: string, content: string) {
    await makeRequest(`/notes/update/${id}`, "PUT", { content })
}

export async function readNoteContent(id: string): Promise<string> {
    return await makeRequest(`/notes/read/${id}`, "GET")
}
