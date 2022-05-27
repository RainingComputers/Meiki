<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { createNote } from "$lib/api/notes"
    import { formatRequestError } from "$lib/api/request"
    import EntryModal from "$cmp/modal/EntryModal.svelte"

    export let error: string = ""

    const dispatchEvent = createEventDispatcher()

    async function create(noteName: string) {
        try {
            const id = await createNote(noteName)
            dispatchEvent("noteCreated", { id })
        } catch {
            error = formatRequestError(error, "creating note")
        }
    }
</script>

<EntryModal
    {error}
    onValue={create}
    buttonLabel="Create note"
    entryLabel="Enter note name"
/>
