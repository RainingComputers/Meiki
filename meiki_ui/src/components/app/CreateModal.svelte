<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { createNote } from "$lib/api/notes"
    import Entry from "$cmp/Entry.svelte"
    import Button from "$cmp/Button.svelte"
    import Error from "$cmp/toast/Error.svelte"

    let entry: Entry
    let error: string = ""

    const dispatchEvent = createEventDispatcher()

    async function create() {
        try {
            await createNote(entry.getValue())
            error = ""
            dispatchEvent("noteCreated")
        } catch {
            error = "Unable to create note, please try again later"
        }
    }
</script>

<div
    class="w-96 max-h-fit m-auto bg-gray-50 border-gray-200 rounded-xl flex flex-col py-5 px-6"
>
    <div class="flex flex-col gap-5">
        {#if error}
            <Error>{error}</Error>
        {/if}
        <Entry label="Enter note name" onEnter={create} bind:this={entry} />
        <Button label="Create note" onClick={create} />
    </div>
</div>
