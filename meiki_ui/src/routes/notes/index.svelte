<script lang="ts">
    import { onMount } from "svelte"
    import { goto } from "$app/navigation"
    import { readNoteContent, updateNote } from "$lib/api/notes"
    import { listNotes, NoteInfo } from "$lib/api/notes"
    import { debounce } from "$lib/utils/debouncer"
    import { onCtrlPlusS } from "$lib/utils/onCtrlPlusS"
    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/toolbar/AppToolbar.svelte"
    import Workbench from "$cmp/app/Workbench.svelte"
    import Modal from "$cmp/modal/Modal.svelte"
    import LogoutModal from "$cmp/app/modal/LogoutModal.svelte"
    import CreateModal from "$cmp/app/modal/CreateModal.svelte"
    import DeleteModal from "$cmp/app/modal/DeleteModal.svelte"

    let workbench: Workbench
    let logoutModal: Modal
    let createModal: Modal
    let deleteModal: Modal

    let currentNote: NoteInfo
    let currentNoteText: string
    let noteList: Array<NoteInfo> = []
    let changesNotSaved: boolean = false

    let explorerActive: boolean = true
    let editorActive: boolean = true
    let rendererActive: boolean = true

    let explorerError: string = ""
    let toolbarError: string = ""

    async function updateNoteList() {
        try {
            noteList = await listNotes()
        } catch {
            explorerError = "Unable to list notes, cannot connect to server"
        }
    }

    async function syncCurrentNote() {
        try {
            if (currentNote) {
                await updateNote(currentNote.id, currentNoteText)
                changesNotSaved = false
            }
        } catch (err) {
            toolbarError = "sync error"
        }
    }

    const debouncedSyncNote = debounce(syncCurrentNote)
    function onTextChange(event: CustomEvent<{ text: string }>) {
        currentNoteText = event.detail.text
        changesNotSaved = true
        debouncedSyncNote()
    }

    async function onFocusAway() {
        await syncCurrentNote()
    }

    async function selectNote(id: string) {
        await onFocusAway()
        try {
            const noteContent = await readNoteContent(id)
            currentNote = { id, title: noteContent.title }
            currentNoteText = noteContent.content
            workbench.setText(noteContent.content)
            editorActive = true
        } catch (err) {
            toolbarError = "read error"
            currentNote = undefined
        }
    }

    async function deselectAllNotes() {
        await onFocusAway()
        currentNote = undefined
    }

    function onNoteCreated(event: CustomEvent<{ id: string }>) {
        const newNoteID: string = event.detail.id
        updateNoteList()
        selectNote(newNoteID)
        editorActive = true
        createModal.closeModal()
    }

    function onNoteDeleted() {
        updateNoteList()
        deselectAllNotes()
        deleteModal.closeModal()
    }

    function toggleExplorer() {
        explorerActive = !explorerActive
    }

    onCtrlPlusS(syncCurrentNote)

    onMount(async () => {
        await updateNoteList()
    })
</script>

<Modal bind:this={logoutModal}>
    <LogoutModal on:loggedOut={() => goto("/login")} />
</Modal>

<Modal bind:this={deleteModal}>
    <DeleteModal
        noteInfo={currentNote}
        on:deleted={onNoteDeleted}
        on:deleteCancelled={() => {
            deleteModal.closeModal()
        }}
    />
</Modal>

<Modal bind:this={createModal}>
    <CreateModal on:noteCreated={onNoteCreated} />
</Modal>

<App>
    <AppToolbar
        title={currentNote?.title || ""}
        showNoteActions={!!currentNote}
        {explorerActive}
        {editorActive}
        {rendererActive}
        {changesNotSaved}
        {toolbarError}
        on:sidebar={toggleExplorer}
        on:edit={() => {
            editorActive = !editorActive
        }}
        on:render={() => {
            rendererActive = !rendererActive
        }}
        on:create={() => {
            createModal.showModal()
        }}
        on:profile={() => {
            logoutModal.showModal()
        }}
        on:delete={() => {
            deleteModal.showModal()
        }}
    />
    <div class="flex flex-row flex-grow w-full">
        {#if explorerActive}
            <AppExplorer
                {noteList}
                error={explorerError}
                selectedNoteID={currentNote?.id}
                on:selectNote={(event) => {
                    selectNote(event.detail.noteID)
                }}
                on:deselectAllNotes={deselectAllNotes}
            />
        {/if}

        <Workbench
            bind:this={workbench}
            showEditorAndRenderer={!!currentNote}
            on:textChange={onTextChange}
            {rendererActive}
            {editorActive}
        />
    </div>
</App>
