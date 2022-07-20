<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import type { NoteInfo } from "$lib/api/notes"
    import Panel from "$cmp/explorer/Panel.svelte"
    import NotesItem from "$cmp/app/NotesItem.svelte"
    import Item from "$cmp/explorer/Item.svelte"
    import Info from "$cmp/explorer/watermark/Info.svelte"
    import ExplorerToolbar from "$cmp/app/toolbar/ExplorerToolbar.svelte"
    import WatermarkError from "$cmp/explorer/watermark/Error.svelte"
    import ToastError from "$cmp/explorer/toast/Error.svelte"

    export let noteList: Array<NoteInfo>
    export let selectedNoteID: string
    export let watermarkError: string = ""
    export let toastError: string = ""

    const dispatchEvent = createEventDispatcher()

    function deselectAllNotes() {
        dispatchEvent("deselectAllNotes")
    }

    function selectNote(id: string) {
        if (selectedNoteID === id) {
            deselectAllNotes()
            return
        }

        dispatchEvent("selectNote", { noteID: id })
    }
</script>

<Panel widthPercentage={25} onClick={deselectAllNotes}>
    <ExplorerToolbar on:create />
    {#if toastError}
        <ToastError message={toastError} />
    {/if}

    {#if !watermarkError}
        {#each noteList as item (item.id)}
            {#if item.id == selectedNoteID}
                <Item
                    onClick={() => {
                        selectNote(item.id)
                    }}
                    checked={true}
                >
                    <NotesItem title={item.title} />
                </Item>
            {:else}
                <Item
                    onClick={() => {
                        selectNote(item.id)
                    }}
                >
                    <NotesItem title={item.title} />
                </Item>
            {/if}
        {/each}

        {#if !noteList.length}
            <Info
                message="Click the 'Create' button on the toolbar to create a new note"
            />
        {/if}
    {:else}
        <WatermarkError message={watermarkError} />
    {/if}
</Panel>
