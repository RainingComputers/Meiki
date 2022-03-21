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
    import { showRenderer, showEditor } from "$cmp/app/Workbench/stores"
</script>

<Toolbar>
    <ToolbarButton checkable={true} name="sidebar" on:sidebar>
        <SidebarIcon />
    </ToolbarButton>
    <span class="px-2" />
    <ToolbarButton
        checkable={true}
        name="editor"
        on:editor={() => {
            $showEditor = !$showEditor
        }}
    >
        <EditIcon />
    </ToolbarButton>
    <ToolbarButton
        checkable={true}
        name="renderer"
        on:renderer={() => {
            $showRenderer = !$showRenderer
        }}
    >
        <PreviewIcon />
    </ToolbarButton>
    <ToolbarButton checkable={false}>
        <TrashIcon />
    </ToolbarButton>
    <ToolbarTitle title="Meiki" />
    <ToolBarAction label="Create" />
    <span class="px-0.5" />
    <ToolbarButton checkable={false} name="shnooshankar" showLabel={true}>
        <UserIcon />
    </ToolbarButton>
</Toolbar>
