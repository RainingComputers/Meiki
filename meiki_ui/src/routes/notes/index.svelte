<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "$lib/api/authStatus"

    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"
    import Workbench from "$cmp/app/userForm/Workbench.svelte"

    let loggedIn: boolean = false
    let showExplorer: boolean = true

    let workbench: Workbench

    onMount(async () => {
        try {
            await authStatus()
            loggedIn = true
        } catch {
            goto(`/login`)
        }
    })

    function toggleExplorer() {
        showExplorer = !showExplorer
    }
</script>

{#if loggedIn}
    <App>
        <AppToolbar
            on:sidebar={toggleExplorer}
            on:editor={() => {
                workbench.toggleEditor()
            }}
            on:renderer={() => {
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
{/if}
