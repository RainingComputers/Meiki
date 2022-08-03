<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
    import { createEventDispatcher } from "svelte"
    const dispatchEvent = createEventDispatcher()
    let queryValue = ""
    const sendQuery = () => {
        dispatchEvent("query", { query: queryValue })
    }
</script>

<div
    class="flex flex-row outline outline-1 focus-within:outline-2 outline-base-2 bg-base-0 focus-within:outline-primary-focus rounded-xl p-2 flex-grow"
>
    <input
        class=" bg-base-0 placeholder-placeholder outline-none border-0 flex-grow text-content"
        placeholder="Search across notes"
        bind:value={queryValue}
        on:input={sendQuery}
    />
    <span class="stroke-placeholder w-6 h-6">
        <SearchIcon />
    </span>
</div>
