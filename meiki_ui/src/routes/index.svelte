<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "$lib/api/authStatus"
    import Toolbar from "$cmp/toolbar/Toolbar.svelte"

    import ToolbarButton from "$cmp/toolbar/Button.svelte"
    import ToolbarTitle from "$cmp/toolbar/Title.svelte"
    import SidebarIcon from "$cmp/icons/SidebarIcon.svelte"
    import EditIcon from "$cmp/icons/EditIcon.svelte"
    import PreviewIcon from "$cmp/icons/PreviewIcon.svelte"
    import TrashIcon from "$cmp/icons/TrashIcon.svelte"
    import UserIcon from "$cmp/icons/UserIcon.svelte"
    import ToolBarAction from "$cmp/toolbar/Action.svelte"
    import Panel from "$cmp/explorer/Panel.svelte"
    import Workbench from "$cmp/explorer/Workbench.svelte"
    import itemList from "$data/itemList.json"
    import Item from "$cmp/explorer/Item.svelte"
    import Root from "$cmp/Root.svelte"

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
    <Root>
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
        <div class="flex flex-row flex-grow w-full">
            <Panel width="22rem">
                {#each itemList as item (item.id)}
                    <Item title={item.title} />
                {/each}
            </Panel>
            <Workbench />
        </div>
    </Root>
{/if}
