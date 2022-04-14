<script lang="ts">
    import { getUsername } from "$lib/api/user"
    import Toolbar from "$cmp/toolbar/Toolbar.svelte"
    import ToolbarButton from "$cmp/toolbar/Button.svelte"
    import ToolbarTitle from "$cmp/toolbar/Title.svelte"
    import SidebarIcon from "$cmp/icons/SidebarIcon.svelte"
    import EditIcon from "$cmp/icons/EditIcon.svelte"
    import PreviewIcon from "$cmp/icons/PreviewIcon.svelte"
    import TrashIcon from "$cmp/icons/TrashIcon.svelte"
    import UserIcon from "$cmp/icons/UserIcon.svelte"
    import ToolBarAction from "$cmp/toolbar/Action.svelte"

    const username = getUsername()
</script>

<Toolbar>
    <ToolbarButton checkable={true} isButtonChecked={true} name="sidebar" on:sidebar>
        <SidebarIcon />
    </ToolbarButton>
    <span class="px-2" />
    <ToolbarButton checkable={true} name="edit" on:edit>
        <EditIcon />
    </ToolbarButton>
    <ToolbarButton checkable={true} name="render" on:render>
        <PreviewIcon />
    </ToolbarButton>
    <ToolbarButton checkable={false}>
        <TrashIcon />
    </ToolbarButton>
    <ToolbarTitle title="Meiki" />
    <ToolBarAction label="Create" />
    <span class="px-0.5" />
    <ToolbarButton checkable={false} name="profile" label={username} on:profile>
        <UserIcon />
    </ToolbarButton>
</Toolbar>
