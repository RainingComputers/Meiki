<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "$lib/api/authStatus"

    import App from "$cmp/App.svelte"
    import AppExplorer from "$cmp/app/AppExplorer.svelte"
    import AppToolbar from "$cmp/app/AppToolbar.svelte"

    let loggedIn: boolean = false

    onMount(async () => {
        try {
            await authStatus()
            loggedIn = true
        } catch {
            goto(`/login`)
        }
    })
</script>

{#if loggedIn}
    <App>
        <AppToolbar />
        <AppExplorer />
    </App>
{/if}
