<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
</script>

<div
    class="flex flex-row border-2 bg-searchbar dark:bg-searchbar-dark border-secondaryFocus focus-within:border-primaryFocus rounded-xl p-2 flex-grow"
>
    <input
        class=" bg-searchbar dark:bg-searchbar-dark placeholder-contentPlaceholder dark:placeholder-contentPlaceholder-dark outline-none border-0 flex-grow"
        placeholder="Don't search yet"
    />
    <span
        class="stroke-contentPlaceholder dark:stroke-contentPlaceholder-dark w-6 h-6"
    >
        <SearchIcon />
    </span>
</div>
