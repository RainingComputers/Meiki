<script lang="ts">
    import { authStatus } from "$lib/api/user"
    import { onMount } from "svelte"
    import { goto } from "$app/navigation"

    // Server side Onload doesn't work because we can't get the session token which is on the client side
    let loggedIn = false
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
    <slot />
{/if}
