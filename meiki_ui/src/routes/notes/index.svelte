<script context="module" lang="ts">
    const NOTE_SYNC_INTERVAL = 5000
</script>

<script lang="ts">
    import { onMount } from "svelte"
    import { goto } from "$app/navigation"
    import { readNoteContent, updateNote } from "$lib/api/notes"
    import { listNotes, NoteInfo } from "$lib/api/notes"
    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"
    import Workbench from "$cmp/app/Workbench.svelte"
    import ModalOverlay from "$cmp/modal/ModalOverlay.svelte"
    import LogoutModal from "$cmp/app/LogoutModal.svelte"
    import CreateModal from "$cmp/app/CreateModal.svelte"
    import DeleteModal from "$cmp/app/DeleteModal.svelte"

    let showExplorer: boolean = true
    let workbench: Workbench
    let logoutModalOverlay: ModalOverlay
    let createModalOverlay: ModalOverlay
    let deleteModalOverlay: ModalOverlay

    let currentNoteID: string
    let currentNoteText: string
    let currentNoteTitle: string
    let noteList: Array<NoteInfo> = []
    let syncIntervalID: ReturnType<typeof setInterval>

    function toggleExplorer() {
        showExplorer = !showExplorer
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
            if (currentNoteID) {
                await updateNote(currentNoteID, currentNoteText)
            }
        } catch {
            // TODO: Error handling, how to show this toast
        }
    }

    function startNoteSync() {
        syncIntervalID = setInterval(async () => {
            await syncCurrentNote()
        }, NOTE_SYNC_INTERVAL)
    }

    function stopNoteSync() {
        if (syncIntervalID) clearInterval(syncIntervalID)
    }

    async function selectNote(id: string) {
        stopNoteSync()

        currentNoteID = id
        currentNoteTitle = "testing"

        try {
            const contentInfo = await readNoteContent(id)
            currentNoteText = contentInfo.content
            currentNoteTitle = contentInfo.title
            workbench.setText(currentNoteText)
            startNoteSync()
        } catch {
            deselectAllNotes()
            // TODO: handle this error
        }
    }

    function deselectAllNotes() {
        stopNoteSync()
        currentNoteID = undefined
        currentNoteText = undefined
        currentNoteTitle = undefined
    }

    function onTextChange(event: CustomEvent<{ text: string }>) {
        currentNoteText = event.detail.text
    }

    function onNoteCreated(event: CustomEvent<{ id: string }>) {
        const newNoteID: string = event.detail.id
        updateNoteList()
        selectNote(newNoteID)
        workbench.enableEditor()
        createModalOverlay.closeModal()
    }

    function onNoteDeleted() {
        updateNoteList()
        deselectAllNotes()
        deleteModalOverlay.closeModal()
    }

    onMount(async () => {
        await updateNoteList()
        startNoteSync()
    })
</script>

<ModalOverlay bind:this={logoutModalOverlay}>
    <LogoutModal on:loggedOut={() => goto("/login")} />
</ModalOverlay>

<ModalOverlay bind:this={deleteModalOverlay}>
    <DeleteModal
        noteID={currentNoteID}
        on:deleted={onNoteDeleted}
        on:deleteCancelled={() => {
            deleteModalOverlay.closeModal()
        }}
    />
</ModalOverlay>

<ModalOverlay bind:this={createModalOverlay}>
    <CreateModal on:noteCreated={onNoteCreated} />
</ModalOverlay>

<App>
    <AppToolbar
        title={currentNoteTitle}
        showNoteActions={!!currentNoteID}
        on:sidebar={toggleExplorer}
        on:edit={() => {
            workbench.toggleEditor()
        }}
        on:render={() => {
            workbench.toggleRenderer()
        }}
        on:create={() => {
            createModalOverlay.showModal()
        }}
        on:profile={() => {
            logoutModalOverlay.showModal()
        }}
        on:delete={() => {
            deleteModalOverlay.showModal()
        }}
    />
    <div class="flex flex-row flex-grow w-full">
        {#if showExplorer}
            <AppExplorer
                {noteList}
                selectedNoteID={currentNoteID}
                on:selectNote={(event) => {
                    selectNote(event.detail.noteID)
                }}
                on:deselectAllNotes={deselectAllNotes}
            />
        {/if}

        <Workbench
            bind:this={workbench}
            showEditorAndRenderer={!!currentNoteID}
            on:textChange={onTextChange}
        />
    </div>
</App>
