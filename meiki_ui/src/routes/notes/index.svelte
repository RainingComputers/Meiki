<script lang="ts">
    import { onMount } from "svelte"
    import { goto } from "$app/navigation"
    import { readNoteContent, renameNote, updateNote } from "$lib/api/notes"
    import { listNotes, NoteInfo } from "$lib/api/notes"
    import { formatRequestError } from "$lib/api/request"
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
    let noteToDelete: NoteInfo
    let noteList: Array<NoteInfo> = []
    let changesNotSaved: boolean = false

    let explorerActive: boolean = true
    let editorActive: boolean = true
    let rendererActive: boolean = true

    let explorerWatermarkError: string = ""
    let explorerToastError: string = ""
    let toolbarError: string = ""

    async function updateNoteList() {
        try {
            noteList = await listNotes()
        } catch (err) {
            explorerWatermarkError = formatRequestError(err, "listing notes")
        }
    }

    async function syncCurrentNote() {
        try {
            if (currentNote) {
                await updateNote(currentNote.id, workbench.getText())
                changesNotSaved = false
            }
        } catch (err) {
            toolbarError = "sync error"
            explorerToastError = formatRequestError(err, "syncing note")
        }
    }

    const debouncedSyncNote = debounce(syncCurrentNote)
    function onTextChange() {
        changesNotSaved = true
        debouncedSyncNote()
    }

    async function selectNote(id: string) {
        await syncCurrentNote()
        try {
            const noteContent = await readNoteContent(id)
            currentNote = { id, title: noteContent.title }
            workbench.setText(noteContent.content)
            editorActive = true
        } catch (err) {
            explorerToastError = formatRequestError(err, "reading note")
            currentNote = undefined
        }
    }

    async function deselectAllNotes() {
        await syncCurrentNote()
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
        if (currentNote.id == noteToDelete.id) currentNote = undefined
        deleteModal.closeModal()
    }

    function toggleExplorer() {
        explorerActive = !explorerActive
    }

    async function onRename(event: CustomEvent<{ newTitle: string }>) {
        try {
            await renameNote(currentNote.id, event.detail.newTitle)
            updateNoteList()
            currentNote.title = event.detail.newTitle
        } catch (err) {
            explorerToastError = formatRequestError(err, "renaming note")
        }
    }

    onCtrlPlusS(syncCurrentNote)

    onMount(async () => {
        await updateNoteList()
    })
</script>

<svelte:head>
    <title>{currentNote?.title || "Meiki"}</title>
</svelte:head>

<App>
    <Modal bind:this={logoutModal}>
        <LogoutModal on:loggedOut={() => goto("/login")} />
    </Modal>

    <Modal bind:this={deleteModal}>
        <DeleteModal
            noteInfo={noteToDelete}
            on:deleted={onNoteDeleted}
            on:deleteCancelled={() => {
                deleteModal.closeModal()
            }}
        />
    </Modal>

    <Modal bind:this={createModal}>
        <CreateModal on:noteCreated={onNoteCreated} />
    </Modal>

    <AppToolbar
        title={currentNote?.title}
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
        on:profile={() => {
            logoutModal.showModal()
        }}
        on:rename={onRename}
    />
    <div class="flex flex-row flex-grow h-full w-full">
        {#if explorerActive}
            <AppExplorer
                on:createNote={() => {
                    createModal.showModal()
                }}
                on:deleteNote={(event) => {
                    noteToDelete = event.detail.item
                    deleteModal.showModal()
                }}
                {noteList}
                watermarkError={explorerWatermarkError}
                toastError={explorerToastError}
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
