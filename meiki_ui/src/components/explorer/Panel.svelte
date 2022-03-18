<script lang="ts">
    import { fly } from "svelte/transition"

    export let width: string = "24rem"
</script>

<div
    class=" bg-gray-50 border-gray-200 border-r flex flex-col gap-1 py-5"
    style="width: {width}"
    data-cy="explorer"
    transition:fly={{ x: -400, duration: 350, opacity: 1 }}
>
    <slot />
</div>
