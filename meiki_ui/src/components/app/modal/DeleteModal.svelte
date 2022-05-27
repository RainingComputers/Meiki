<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { deleteNote, NoteInfo } from "$lib/api/notes"
    import { formatRequestError } from "$lib/api/request"
    import ConfirmModal from "$cmp/modal/ConfirmModal.svelte"

    export let noteInfo: NoteInfo
    export let error: string = ""

    const dispatchEvent = createEventDispatcher()

    async function deleteCurrentNote() {
        try {
            await deleteNote(noteInfo.id)
            dispatchEvent("deleted")
        } catch (err) {
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
    onYes={deleteCurrentNote}
    onNo={cancelDelete}
/>
