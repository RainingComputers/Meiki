<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { deleteNote } from "$lib/api/notes"
    import ConfirmModal from "$cmp/modal/ConfirmModal.svelte"

    export let noteID: string
    export let error: string = ""
    const dispatchEvent = createEventDispatcher()

    async function deleteCurrentNote() {
        try {
            await deleteNote(noteID)
            dispatchEvent("deleted")
        } catch {
            // TODO: Error handling
        }
    }

    async function cancelDelete() {
        dispatchEvent("deleteCancelled")
    }
</script>

<!-- TODO: note title here in the are you sure message -->
<ConfirmModal
    {error}
    message={`Are you sure you want to delete ${noteID}?`}
    onYes={deleteCurrentNote}
    onNo={cancelDelete}
/>
