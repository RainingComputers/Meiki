import { makeRequest } from "./request"

export type NoteInfo = {
    id: string
    title: string
}

export type NoteContentInfo = {
    title: string
    content: string
}

export async function listNotes(): Promise<Array<NoteInfo>> {
    return await makeRequest("/notes/list", "GET")
}

export async function searchNotes(query: string): Promise<Array<NoteInfo>> {
    const response = await makeRequest(`/notes/search?query=${query}`, "GET")
    console.log(JSON.stringify(response))
    return response
}

export async function createNote(title: string): Promise<string> {
    return await makeRequest("/notes/create", "POST", { title })
}

export async function deleteNote(id: string) {
    await makeRequest(`/notes/delete/${id}`, "DELETE")
}

export async function updateNote(id: string, content: string) {
    await makeRequest(`/notes/update/${id}`, "PUT", { content })
}

export async function readNoteContent(id: string): Promise<NoteContentInfo> {
    return await makeRequest(`/notes/read/${id}`, "GET")
}

export async function renameNote(id: string, newTitle: string) {
    await makeRequest(`/notes/rename/${id}`, "PUT", { title: newTitle })
}
