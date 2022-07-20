<script lang="ts">
    export let onClick = () => {}
    export let checked = false
    $: checkedClass = checked
        ? "bg-blue-500 text-white stroke-white"
        : " text-slate-900 stroke-slate-900"
</script>

<div
    class=" {checkedClass} hover:bg-blue-500 hover:text-white hover:stroke-white "
    on:click|stopPropagation={onClick}
>
    <slot />
</div>
