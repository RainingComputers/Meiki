<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { createNote } from "$lib/api/notes"
    import Entry from "$cmp/Entry.svelte"
    import Button from "../Button.svelte"

    let entry: Entry

    const dispatchEvent = createEventDispatcher()

    async function create() {
        // TODO: handle errors

        await createNote(entry.getValue())
        dispatchEvent("noteCreated")
    }
</script>

<div
    class="w-96 max-h-fit m-auto bg-gray-50 border-gray-200 rounded-xl flex flex-col py-5 px-6"
>
    <div class="flex flex-col gap-5">
        <Entry
            label="Enter note name"
            onEnter={create}
            bind:this={entry}
        />
        <Button label="Create note" onClick={create} />
    </div>
</div>
