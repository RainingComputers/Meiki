<script lang="ts">
    import Panel from "$cmp/explorer/Panel.svelte"
    import Workbench from "$cmp/app/Workbench.svelte"
    import itemList from "$data/itemList.json"
    import Item from "$cmp/explorer/Item.svelte"
</script>

<div class="flex flex-row flex-grow w-full">
    <Panel width="22rem">
        {#each itemList as item (item.id)}
            <Item title={item.title} />
        {/each}
    </Panel>
    <Workbench />
</div>
