<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { deleteNote } from "$lib/api/notes"
    import currentNote from "$lib/stores/currentNote"
    import ConfirmModal from "$cmp/modal/ConfirmModal.svelte"

    export let error: string = ""
    const dispatchEvent = createEventDispatcher()

    async function deleteCurrentNote() {
        try {
            await deleteNote($currentNote)
            dispatchEvent("deleted")
        } catch {
            // TODO: Error handling
            // TODO: make this a modal
        }
    }

    async function cancelDelete() {
        dispatchEvent("deleteCancelled")
    }
</script>

<ConfirmModal
    {error}
    message={`Are you sure you want to delete ${$currentNote}`}
    onYes={deleteCurrentNote}
    onNo={cancelDelete}
/>
