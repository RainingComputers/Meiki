<script context="module" lang="ts">
    const NOTE_SYNC_INTERVAL = 5000
</script>

<script lang="ts">
    import { onMount } from "svelte"
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

    let currentNoteID: string = undefined
    let currentNoteText: string = undefined
    let noteList: Array<NoteInfo> = []
    let syncIntervalID: ReturnType<typeof setInterval> = undefined

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

        try {
            currentNoteText = await readNoteContent(id)
            workbench.setText(currentNoteText)
        } catch {
            currentNoteText = undefined
            // TODO: handle this error
        }

        startNoteSync()
    }

    function deselectAllNotes() {
        stopNoteSync()
        currentNoteID = undefined
        currentNoteText = undefined
    }

    function onTextChange(event: any) {
        currentNoteText = event.detail.text
    }

    function onNoteCreated(event: any) {
        const noteInfo: NoteInfo = event.detail.noteInfo
        updateNoteList()
        selectNote(noteInfo.id)
        workbench.enableEditor()
        createModalOverlay.closeModal()
    }

    function onNoteDeleted(event: any) {
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
    <LogoutModal />
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
