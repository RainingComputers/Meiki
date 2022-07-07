<script lang="ts">
    import { fly } from "svelte/transition"

    export let widthPercentage: number = 25
    let width: number = (widthPercentage * window.innerWidth) / 100
    export let onClick: () => void = undefined

    let expanding = false

    function startExpand(event: MouseEvent) {
        expanding = true
        width = event.pageX
    }

    function stopExpand() {
        expanding = false
    }

    function expand(event: MouseEvent) {
        if (!expanding) return
        width = event.pageX
    }
</script>

<svelte:window on:mouseup={stopExpand} on:mousemove={expand} />

<div
    class=" bg-gray-50 border-gray-200 flex flex-col gap-1 h-full overflow-y-scroll"
    style="width: {width}px"
    data-cy="explorer"
    on:click={onClick}
    transition:fly|local={{ x: -400, duration: 200, opacity: 1 }}
>
    <slot />
    <span class="py-10" />
</div>
<div
    data-cy="explorer-expander"
    class=" cursor-col-resize h-full border-blue-600 border-2"
    on:mousedown={startExpand}
    on:mousemove={expand}
/>

<style>
</style>
