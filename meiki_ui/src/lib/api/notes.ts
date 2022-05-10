import { makeRequest } from "./request"
import { debounce } from "$lib/utils/debouncer"

export type NoteInfo = {
    id: string
    title: string
    username: string
}

export async function listNotes(): Promise<Array<NoteInfo>> {
    return await makeRequest("/notes/list", "GET")
}

export async function createNote(title: string) {
    await makeRequest("/notes/create", "POST", { title })
}

export async function deleteNote(id: string) {
    await makeRequest(`/notes/delete/${id}`, "DELETE")
}

export async function updateNote(id: string, content: string) {
    await makeRequest(`/notes/update/${id}`, "PUT", { content })
}

export const debouncedUpdateNote = debounce(updateNote)
