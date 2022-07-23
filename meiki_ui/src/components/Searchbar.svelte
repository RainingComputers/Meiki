<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
</script>

<div
    class="flex flex-row border-2 bg-white border-gray-200 focus-within:border-blue-500 rounded-xl p-2 flex-grow min-w-0"
>
    <input
        class=" bg-white placeholder-gray-400 outline-none border-0 flex-grow min-w-0"
        placeholder="Search"
        on:change
    />
    <span class="stroke-gray-400 w-6 h-6">
        <SearchIcon />
    </span>
</div>
