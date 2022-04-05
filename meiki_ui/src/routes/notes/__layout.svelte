<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"
    import { authStatus } from "$lib/api/user"

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
    <slot />
{/if}
