<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import Searchbar from "$cmp/Searchbar.svelte"
    import Button from "$cmp/Button.svelte"

    const dispatchEvent = createEventDispatcher()
</script>

<div class="flex flex-row px-4 py-3 gap-2">
    <Searchbar on:query />
    <Button label="+" onClick={() => dispatchEvent("createNote")} />
</div>
