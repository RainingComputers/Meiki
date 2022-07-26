<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
</script>

<div
    class="flex flex-row border-2 bg-white border-background-secondary-focus focus-within:border-primary-content rounded-xl p-2 flex-grow"
>
    <input
        class=" bg-white placeholder-content-placeholder outline-none border-0 flex-grow"
        placeholder="Don't search yet"
    />
    <span class="stroke-content-placeholder w-6 h-6">
        <SearchIcon />
    </span>
</div>
