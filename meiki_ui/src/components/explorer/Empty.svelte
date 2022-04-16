<script lang="ts">
    import ItemsIcon from "$cmp/icons/ItemsIcon.svelte"

    export let message: string
</script>

<div
    class="flex flex-col gap-5 h-full items-center justify-center text-slate-900 stroke-slate-900 opacity-20"
>
    <span class="flex-shrink-0 w-20 h-20">
        <ItemsIcon />
    </span>
    <span class="px-10 font-bold text-xl text-center">
        {message}
    </span>
</div>
