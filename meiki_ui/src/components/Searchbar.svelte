<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
</script>

<div
    class="flex flex-row border-2 bg-searchbar border-secondaryFocus focus-within:border-primaryFocus rounded-xl p-2 flex-grow"
>
    <input
        class=" bg-white placeholder-contentPlaceholder outline-none border-0 flex-grow"
        placeholder="Don't search yet"
    />
    <span class="stroke-contentPlaceholder w-6 h-6">
        <SearchIcon />
    </span>
</div>
