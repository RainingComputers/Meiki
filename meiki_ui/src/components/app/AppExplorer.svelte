<script lang="ts">
    import Panel from "$cmp/explorer/Panel.svelte"
    import itemList from "$data/itemList.json"
    import Item from "$cmp/explorer/Item.svelte"
</script>

<Panel width="20%">
    {#each itemList as item (item.id)}
        <Item title={item.title} />
    {/each}
</Panel>
