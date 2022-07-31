<script lang="ts">
    import SearchIcon from "$cmp/icons/SearchIcon.svelte"
</script>

<div
    class="flex flex-row border-2 bg-base-0 border-base-2 focus-within:border-primary-focus rounded-xl p-2 flex-grow"
>
    <input
        class=" bg-white placeholder-base-4 outline-none border-0 flex-grow"
        placeholder="Don't search yet"
    />
    <span class="stroke-base-4 w-6 h-6">
        <SearchIcon />
    </span>
</div>
