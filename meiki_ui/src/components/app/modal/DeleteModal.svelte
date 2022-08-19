<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { deleteNote, type NoteInfo } from "$lib/api/notes"
    import { formatRequestError } from "$lib/api/request"
    import ConfirmModal from "$cmp/modal/ConfirmModal.svelte"

    export let noteInfo: NoteInfo
    export let error: string = ""

    const dispatchEvent = createEventDispatcher()

    async function deleteSelectedNote() {
        try {
            await deleteNote(noteInfo.id)
            dispatchEvent("deleted")
        } catch (err) {
            console.log(err)
            error = formatRequestError(err, "deleting note")
        }
    }

    async function cancelDelete() {
        dispatchEvent("deleteCancelled")
    }
</script>

<ConfirmModal
    {error}
    message={`Are you sure you want to delete ${noteInfo.title}?`}
    onYes={deleteSelectedNote}
    onNo={cancelDelete}
/>
