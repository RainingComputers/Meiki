<script lang="ts">
    export let onClick = () => {}
    export let checked = false
    $: checkedClass = checked
        ? "bg-primaryFocus text-contentPrimaryFocus stroke-contentPrimaryFocus dark:text-contentPrimaryFocus-dark dark:stroke-contentPrimaryFocus-dark"
        : "text-info text-contentPrimary dark:text-contentPrimary-dark stroke-contentPrimary dark:stroke-contentPrimary-dark stroke-info"
</script>

<div
    class=" {checkedClass} hover:bg-primaryFocus hover:text-contentPrimaryFocus hover:stroke-contentPrimaryFocus dark:hover:text-contentPrimaryFocus-dark dark:hover:stroke-contentPrimaryFocus-dark"
    on:click|stopPropagation={onClick}
>
    <slot />
</div>
