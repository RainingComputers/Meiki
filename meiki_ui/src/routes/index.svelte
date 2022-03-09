<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "../api/authStatus"
    import Toolbar from "../components/toolbar/Toolbar.svelte"

    import ToolbarButton from "../components/toolbar/Button.svelte"
    import ToolbarTitle from "../components/toolbar/Title.svelte"
    import SidebarIcon from "../components/icons/SidebarIcon.svelte"
    import EditIcon from "../components/icons/EditIcon.svelte"
    import PreviewIcon from "../components/icons/PreviewIcon.svelte"
    import TrashIcon from "../components/icons/TrashIcon.svelte"
    import UserIcon from "../components/icons/UserIcon.svelte"
    import ToolBarAction from "../components/toolbar/Action.svelte"

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
        <ToolbarButton checked={true}>
            <SidebarIcon />
        </ToolbarButton>
        <span class="px-2" />
        <ToolbarButton checked={true}>
            <EditIcon />
        </ToolbarButton>
        <ToolbarButton checked={true}>
            <PreviewIcon />
        </ToolbarButton>
        <ToolbarButton checked={false}>
            <TrashIcon />
        </ToolbarButton>
        <ToolbarTitle title="Meiki" />
        <ToolBarAction label="Create" />
        <span class="px-0.5" />
        <ToolbarButton checked={false} label="shnooshankar">
            <UserIcon />
        </ToolbarButton>
    </Toolbar>
{/if}
