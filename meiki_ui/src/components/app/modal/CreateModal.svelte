<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { createNote } from "$lib/api/notes"
    import EntryModal from "$cmp/modal/EntryModal.svelte"

    export let error: string = ""

    const dispatchEvent = createEventDispatcher()

    async function create(noteName: string) {
        try {
            const id = await createNote(noteName)
            dispatchEvent("noteCreated", { id })
        } catch {
            error = "Unable to create note, unable to connect to server"
        }
    }
</script>

<EntryModal
    {error}
    onValue={create}
    buttonLabel="Create note"
    entryLabel="Enter note name"
/>
