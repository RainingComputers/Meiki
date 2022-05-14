<script lang="ts">
    import { onMount } from "svelte"
    import { updateNote } from "$lib/api/notes"
    import currentNote from "$lib/stores/currentNote"
    import currentNoteText from "$lib/stores/currentNoteText"
    import { debounce } from "$lib/utils/debouncer"
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
    let explorer: AppExplorer

    export const debouncedUpdateNote = debounce(updateNote)

    function toggleExplorer() {
        showExplorer = !showExplorer
    }

    async function syncCurrentNote() {
        try {
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

<ModalOverlay bind:this={deleteModalOverlay}>
    <DeleteModal
        on:deleted={() => {
            explorer.updateItems()
            deleteModalOverlay.closeModal()
        }}
        on:deleteCancelled={() => {
            deleteModalOverlay.closeModal()
        }}
    />
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
        showNoteActions={!!$currentNote}
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
            <AppExplorer bind:this={explorer} />
        {/if}

        <Workbench showRenderer={true} bind:this={workbench} />
    </div>
</App>
