<script lang="ts">
    import { onMount } from "svelte"
    import { listNotes, NoteInfo } from "$lib/api/notes"
    import Panel from "$cmp/explorer/Panel.svelte"
    import Item from "$cmp/explorer/Item.svelte"
    import Empty from "$cmp/explorer/Empty.svelte"

    let itemList: Array<NoteInfo> = []

    onMount(async () => {
        itemList = await listNotes()
    })
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
