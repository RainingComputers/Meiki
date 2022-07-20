<script lang="ts">
    import ItemsIcon from "$cmp/icons/ItemsIcon.svelte"
    export let title: string
</script>

<div class="flex flex-row px-4 gap-3 py-1 items-center text-inherit">
    <span class="flex-shrink-0 w-5 h-5">
        <ItemsIcon />
    </span>
    <span>{title}</span>
</div>
