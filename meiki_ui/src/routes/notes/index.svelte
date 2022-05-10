<script lang="ts">
    import { onMount } from "svelte"
    import currentNote from "$lib/stores/currentNote"
    import currentNoteText from "$lib/stores/currentNoteText"
    import { deleteNote, debouncedUpdateNote } from "$lib/api/notes"

    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"
    import Workbench from "$cmp/app/Workbench.svelte"
    import ModalOverlay from "$cmp/ModalOverlay.svelte"
    import LogoutModal from "$cmp/app/LogoutModal.svelte"
    import CreateModal from "$cmp/app/CreateModal.svelte"

    let showExplorer: boolean = true

    let workbench: Workbench
    let logoutModalOverlay: ModalOverlay
    let createModalOverlay: ModalOverlay
    let explorer: AppExplorer

    function toggleExplorer() {
        showExplorer = !showExplorer
    }

    async function deleteCurrentNote() {
        try {
            await deleteNote($currentNote)
            explorer.updateItems()
        } catch {
            // TODO: Error handling
            // TODO: make this a modal
        }
    }

    async function syncCurrentNote() {
        try {
            console.log("Syncing")
            await debouncedUpdateNote($currentNote, $currentNoteText)
        } catch {
            // TODO: Error handling, how to show this toast
        }
    }

    onMount(() => {
        setInterval(async () => {
            await syncCurrentNote()
        }, 5000)
    })
</script>

<ModalOverlay bind:this={logoutModalOverlay}>
    <LogoutModal />
</ModalOverlay>

<ModalOverlay bind:this={createModalOverlay}>
    <CreateModal
        on:noteCreated={() => {
            explorer.updateItems()
            createModalOverlay.closeModal()
        }}
    />
</ModalOverlay>

<App>
    <AppToolbar
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
            deleteCurrentNote()
        }}
    />
    <div class="flex flex-row flex-grow w-full">
        {#if showExplorer}
            <AppExplorer bind:this={explorer} />
        {/if}
        <Workbench bind:this={workbench} />
    </div>
</App>
