<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import Searchbar from "$cmp/Searchbar.svelte"
    import Action from "$cmp/toolbar/Action.svelte"

    const dispatchEvent = createEventDispatcher()
</script>

<div class="flex flex-row px-4 py-3 gap-2">
    <Searchbar />
    <Action label="+" onClick={() => dispatchEvent("createNote")} />
</div>
