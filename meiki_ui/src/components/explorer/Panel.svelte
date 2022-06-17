<script lang="ts">
    import { fly } from "svelte/transition"

    export let width: string = "24rem"
    export let onClick: () => void = undefined
</script>

<div
    class=" bg-gray-50 border-gray-200 border-r flex flex-col gap-1 h-full overflow-y-scroll"
    style="width: {width}"
    data-cy="explorer"
    on:click={onClick}
    transition:fly|local={{ x: -400, duration: 200, opacity: 1 }}
>
    <slot />
    <span class="py-10" />
</div>
