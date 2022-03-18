<script lang="ts">
    import Toolbar from "$cmp/toolbar/Toolbar.svelte"

    import ToolbarButton from "$cmp/toolbar/Button.svelte"
    import ToolbarTitle from "$cmp/toolbar/Title.svelte"
    import SidebarIcon from "$cmp/icons/SidebarIcon.svelte"
    import EditIcon from "$cmp/icons/EditIcon.svelte"
    import PreviewIcon from "$cmp/icons/PreviewIcon.svelte"
    import TrashIcon from "$cmp/icons/TrashIcon.svelte"
    import UserIcon from "$cmp/icons/UserIcon.svelte"
    import ToolBarAction from "$cmp/toolbar/Action.svelte"
</script>

<Toolbar>
    <ToolbarButton checked={true} name="sidebar" on:sidebar>
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
    <ToolbarButton checked={false} name="shnooshankar" showLabel={true}>
        <UserIcon />
    </ToolbarButton>
</Toolbar>
