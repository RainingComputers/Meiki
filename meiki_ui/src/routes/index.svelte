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
    import Panel from "../components/explorer/Panel.svelte"
    import Workbench from "../components/explorer/Workbench.svelte"
    import itemList from "../itemList.json"
    import Item from "../components/explorer/Item.svelte"

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
    <section class="flex flex-col items-stretch h-screen">
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
        <div class="flex flex-row flex-grow">
            <Panel width="22rem">
                {#each itemList as item (item.id)}
                    <Item title={item.title} />
                {/each}
            </Panel>
            <Workbench />
        </div>
    </section>
{/if}
