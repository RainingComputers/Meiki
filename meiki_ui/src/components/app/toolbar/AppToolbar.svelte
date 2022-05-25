<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { getUsername } from "$lib/api/user"
    import Toolbar from "$cmp/toolbar/Toolbar.svelte"
    import ToolbarButton from "$cmp/toolbar/Button.svelte"
    import ToolbarTitle from "$cmp/toolbar/Title.svelte"
    import ToolbarGroup from "$cmp/toolbar/Group.svelte"
    import SidebarIcon from "$cmp/icons/SidebarIcon.svelte"
    import EditIcon from "$cmp/icons/EditIcon.svelte"
    import PreviewIcon from "$cmp/icons/PreviewIcon.svelte"
    import TrashIcon from "$cmp/icons/TrashIcon.svelte"
    import UserIcon from "$cmp/icons/UserIcon.svelte"
    import ToolBarAction from "$cmp/toolbar/Action.svelte"
    import ToolbarSyncIndicator from "$cmp/app/toolbar/SyncIndicator.svelte"

    export let showNoteActions: boolean
    export let title: string

    export let explorerActive: boolean
    export let editorActive: boolean
    export let rendererActive: boolean

    export let changesNotSaved: boolean

    const username = getUsername()
    const dispatchEvent = createEventDispatcher()
</script>

<Toolbar>
    <ToolbarButton isButtonChecked={explorerActive} name="sidebar" on:sidebar>
        <SidebarIcon />
    </ToolbarButton>

    <ToolbarGroup show={showNoteActions}>
        <span class="px-2" />
        <ToolbarButton isButtonChecked={editorActive} name="edit" on:edit>
            <EditIcon />
        </ToolbarButton>
        <ToolbarButton isButtonChecked={rendererActive} name="render" on:render>
            <PreviewIcon />
        </ToolbarButton>
        <ToolbarButton name="delete" on:delete>
            <TrashIcon />
        </ToolbarButton>
    </ToolbarGroup>

    <ToolbarTitle {title}>
        <ToolbarSyncIndicator show={showNoteActions} {changesNotSaved} />
    </ToolbarTitle>

    <ToolBarAction label="Create" onClick={() => dispatchEvent("create")} />
    <span class="px-0.5" />
    <ToolbarButton name="profile" label={username} on:profile>
        <UserIcon />
    </ToolbarButton>
</Toolbar>
