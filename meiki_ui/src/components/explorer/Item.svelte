<script lang="ts">
    export let onClick = () => {}
    export let checked = false
    $: checkedClass = checked
        ? "bg-primaryFocus text-contentPrimary stroke-contentPrimary"
        : "text-info stroke-info"
</script>

<div
    class=" {checkedClass} hover:bg-primaryFocus hover:text-contentPrimary hover:stroke-contentPrimary "
    on:click|stopPropagation={onClick}
>
    <slot />
</div>
