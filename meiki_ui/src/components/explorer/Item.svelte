<script lang="ts">
    export let onClick = () => {}
    export let checked = false
    $: checkedClass = checked
        ? "bg-primary-focus text-primary-content stroke-primary-content"
        : "text-content stroke-content"
</script>

<div
    class=" {checkedClass} hover:bg-primary-focus hover:text-primary-content hover:stroke-primary-content "
    on:click|stopPropagation={onClick}
>
    <slot />
</div>
