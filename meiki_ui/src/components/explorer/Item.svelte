<script lang="ts">
    export let onClick = () => {}
    export let checked = false
    $: checkedClass = checked
        ? "bg-background-primary-focus text-content-primary stroke-content-primary"
        : "text-info stroke-info"
</script>

<div
    class=" {checkedClass} hover:bg-background-primary-focus hover:text-content-primary hover:stroke-content-primary "
    on:click|stopPropagation={onClick}
>
    <slot />
</div>
