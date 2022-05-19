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

    export let showNoteActions: boolean = false
    export let title: string

    const username = getUsername()
    const dispatchEvent = createEventDispatcher()
</script>

<Toolbar>
    <ToolbarButton
        checkable={true}
        isButtonChecked={true}
        name="sidebar"
        on:sidebar
    >
        <SidebarIcon />
    </ToolbarButton>

    <ToolbarGroup show={showNoteActions}>
        <span class="px-2" />
        <ToolbarButton
            checkable={true}
            isButtonChecked={true}
            name="edit"
            on:edit
        >
            <EditIcon />
        </ToolbarButton>
        <ToolbarButton
            checkable={true}
            isButtonChecked={true}
            name="render"
            on:render
        >
            <PreviewIcon />
        </ToolbarButton>
        <ToolbarButton checkable={false} name="delete" on:delete>
            <TrashIcon />
        </ToolbarButton>
    </ToolbarGroup>

    <ToolbarTitle title={title || ""} />

    <ToolBarAction label="Create" onClick={() => dispatchEvent("create")} />
    <span class="px-0.5" />
    <ToolbarButton checkable={false} name="profile" label={username} on:profile>
        <UserIcon />
    </ToolbarButton>
</Toolbar>
