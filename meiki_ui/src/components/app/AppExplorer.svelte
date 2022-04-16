<script lang="ts">
    import itemList from "$data/itemList.json"
    import Panel from "$cmp/explorer/Panel.svelte"
    import Item from "$cmp/explorer/Item.svelte"
    import Empty from "$cmp/explorer/Empty.svelte"
</script>

<Panel width="20%">
    {#each itemList as item (item.id)}
        <Item title={item.title} />
    {/each}

    {#if !itemList.length}
        <!-- TODO: Test this empty message in cypress -->
        <Empty
            message="Click the 'Create' button on the toolbar to create a new note"
        />
    {/if}
</Panel>
