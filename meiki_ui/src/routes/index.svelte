<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "../api/authStatus"
    import Toolbar from "../components/toolbar/Toolbar.svelte"

    import ToolbarButton from "../components/toolbar/ToolbarButton.svelte"
    import SidebarIcon from "../components/icons/SidebarIcon.svelte"

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
    <Toolbar>
        <ToolbarButton>
            <SidebarIcon />
        </ToolbarButton>
    </Toolbar>
{/if}
