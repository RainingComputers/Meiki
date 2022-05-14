<script lang="ts">
    import ItemsIcon from "$cmp/icons/ItemsIcon.svelte"

    export let title: string
    export let checked: boolean = false
    export let onClick: any

    $: checkedClass = checked ? "bg-blue-500 text-white stroke-white" : ""
</script>

<div
    class=" text-slate-900 stroke-slate-900 hover:bg-blue-500 hover:text-white hover:stroke-white"
    on:click|stopPropagation={onClick}
>
    <div class="flex flex-row px-4 gap-3 py-1 items-center {checkedClass}">
        <span class="flex-shrink-0 w-5 h-5">
            <ItemsIcon />
        </span>
        <span>{title}</span>
    </div>
</div>
