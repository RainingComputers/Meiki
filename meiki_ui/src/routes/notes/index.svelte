<script context="module" lang="ts">
    const NOTE_SYNC_INTERVAL = 5000
</script>

<script lang="ts">
    import { onMount } from "svelte"
    import { goto } from "$app/navigation"
    import { readNoteContent, updateNote } from "$lib/api/notes"
    import { listNotes, NoteInfo } from "$lib/api/notes"
    import { debounce } from "$lib/utils/debouncer"
    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"
    import Workbench from "$cmp/app/Workbench.svelte"
    import Modal from "$cmp/modal/Modal.svelte"
    import LogoutModal from "$cmp/app/LogoutModal.svelte"
    import CreateModal from "$cmp/app/CreateModal.svelte"
    import DeleteModal from "$cmp/app/DeleteModal.svelte"

    let workbench: Workbench
    let logoutModal: Modal
    let createModal: Modal
    let deleteModal: Modal

    let currentNote: NoteInfo
    let currentNoteText: string
    let noteList: Array<NoteInfo> = []
    let lastSavedTime: Date

    let explorerActive: boolean = true
    let editorActive: boolean = true
    let rendererActive: boolean = true

    function toggleExplorer() {
        explorerActive = !explorerActive
    }

    async function updateNoteList() {
        try {
            noteList = await listNotes()
        } catch {
            // TODO: handle this error
        }
    }

    async function syncCurrentNote() {
        try {
            if (currentNote) {
                await updateNote(currentNote.id, currentNoteText)
                lastSavedTime = new Date()
            }
        } catch (err) {
            console.log(err)
            // TODO: Error handling, how to show this toast
        }
    }

    const debouncedSyncNote = debounce(syncCurrentNote)
    function onTextChange(event: CustomEvent<{ text: string }>) {
        currentNoteText = event.detail.text
        debouncedSyncNote()
    }

    async function onFocusAway() {
        await syncCurrentNote()
    }

    async function selectNote(id: string) {
        await onFocusAway()
        try {
            editorActive = true
            const noteContent = await readNoteContent(id)
            currentNote = { id, title: noteContent.title }
            currentNoteText = noteContent.content
            workbench.setText(noteContent.content)
        } catch (err) {
            console.log(err)
            deselectAllNotes()
            // TODO: handle this error
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

    document.addEventListener(
        "keydown",
        function (e) {
            if (
                e.key === "s" &&
                (navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey)
            ) {
                e.preventDefault()
                syncCurrentNote()
            }
        },
        false
    )
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
        {lastSavedTime}
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
