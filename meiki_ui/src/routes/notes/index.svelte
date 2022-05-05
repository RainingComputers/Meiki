<script lang="ts">
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

    function toggleExplorer() {
        showExplorer = !showExplorer
    }
</script>

<ModalOverlay bind:this={logoutModalOverlay}>
    <LogoutModal />
</ModalOverlay>

<ModalOverlay bind:this={createModalOverlay}>
    <CreateModal
        on:noteCreated={() => {
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
    />
    <div class="flex flex-row flex-grow w-full">
        {#if showExplorer}
            <AppExplorer />
        {/if}
        <Workbench bind:this={workbench} />
    </div>
</App>
