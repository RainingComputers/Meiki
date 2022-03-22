<script lang="ts">
    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"
    import Workbench from "$cmp/app/userForm/Workbench.svelte"

    let showExplorer: boolean = true

    let workbench: Workbench

    function toggleExplorer() {
        showExplorer = !showExplorer
    }
</script>

<App>
    <AppToolbar
        on:sidebar={toggleExplorer}
        on:edit={() => {
            workbench.toggleEditor()
        }}
        on:render={() => {
            workbench.toggleRenderer()
        }}
    />
    <div class="flex flex-row flex-grow w-full">
        {#if showExplorer}
            <AppExplorer />
        {/if}
        <Workbench bind:this={workbench} />
    </div>
</App>
