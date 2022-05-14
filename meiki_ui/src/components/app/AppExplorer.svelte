<script lang="ts">
    import { onMount } from "svelte"
    import { listNotes, NoteInfo, readNoteContent } from "$lib/api/notes"
    import currentNoteText from "$lib/stores/currentNoteText"
    import currentNote from "$lib/stores/currentNote"
    import Panel from "$cmp/explorer/Panel.svelte"
    import Item from "$cmp/explorer/Item.svelte"
    import Empty from "$cmp/explorer/Empty.svelte"

    let itemList: Array<NoteInfo> = []

    onMount(async () => {
        await updateItems()
    })

    export async function updateItems() {
        try {
            itemList = await listNotes()
        } catch {
            // TODO: handle this error
        }
    }

    function delselectAllNotes() {
        currentNote.set("")
    }

    async function selectNote(id: string) {
        if ($currentNote == id) {
            delselectAllNotes()
            return
        }

        currentNote.set(id)

        try {
            const noteContent = await readNoteContent(id)
            currentNoteText.set(noteContent)
        } catch {
            // TODO: handle this error
        }
    }
</script>

<Panel width="20%" onClick={delselectAllNotes}>
    {#each itemList as item (item.id)}
        {#if item.id == $currentNote}
            <Item
                title={item.title}
                checked={true}
                onClick={() => {
                    selectNote(item.id)
                }}
            />
        {:else}
            <Item
                title={item.title}
                onClick={() => {
                    selectNote(item.id)
                }}
            />
        {/if}
    {/each}

    {#if !itemList.length}
        <!-- TODO: Test this empty message in cypress -->
        <Empty
            message="Click the 'Create' button on the toolbar to create a new note"
        />
    {/if}
</Panel>
