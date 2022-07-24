<script lang="ts">
    import ItemsIcon from "$cmp/icons/ItemsIcon.svelte"
    import TrashIcon from "$cmp/icons/TrashIcon.svelte"

    export let title: string
    export let onTrashClick = () => {}

    let trashInvisible = true

    const showTrash = () => {
        trashInvisible = false
    }

    const hideTrash = () => {
        trashInvisible = true
    }
</script>

<div
    class="flex flex-row px-4 gap-3 py-1 items-center text-inherit"
    on:mouseenter={showTrash}
    on:mouseleave={hideTrash}
>
    <span class="flex-shrink-0 w-5 h-5">
        <ItemsIcon />
    </span>

    <span class="flex-grow">{title}</span>

    <span
        class="flex-shrink-0 w-5 h-5 cursor-pointer"
        class:hidden={trashInvisible}
        on:click|stopPropagation={onTrashClick}
    >
        <TrashIcon />
    </span>
</div>
