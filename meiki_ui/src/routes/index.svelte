<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "../api/authStatus"
    import Toolbar from "../components/toolbar/Toolbar.svelte"

    import ToolbarButton from "../components/toolbar/ToolbarButton.svelte"
    import SidebarIcon from "../components/icons/SidebarIcon.svelte"
    import EditIcon from "../components/icons/EditIcon.svelte"
    import PreviewIcon from "../components/icons/PreviewIcon.svelte"
    import TrashIcon from "../components/icons/TrashIcon.svelte"

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
        <ToolbarButton isCheckedButton={true}>
            <SidebarIcon />
        </ToolbarButton>
        <span class="px-2" />
        <ToolbarButton isCheckedButton={true}>
            <EditIcon />
        </ToolbarButton>
        <ToolbarButton isCheckedButton={true}>
            <PreviewIcon />
        </ToolbarButton>
        <ToolbarButton isCheckedButton={false}>
            <TrashIcon />
        </ToolbarButton>
    </Toolbar>
{/if}
