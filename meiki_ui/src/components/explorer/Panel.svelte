<script lang="ts">
    import { fly } from "svelte/transition"
    import { createEventDispatcher, onMount } from "svelte"

    const STORE_KEY = "explorer-width"

    export let defaultWidthPercentage: number = 25
    export let onClick: () => void = undefined

    function getInitialWidth() {
        const defaultWidth = (defaultWidthPercentage * window.innerWidth) / 100
        const storedWidth = parseFloat(localStorage.getItem(STORE_KEY))
        return storedWidth || defaultWidth
    }

    let width: number = getInitialWidth()
    let expanding = false

    const dispatch = createEventDispatcher()

    function dispatchResize() {
        dispatch("resize", { width: `${width}px` })
    }

    function startExpand(event: MouseEvent) {
        expanding = true
        width = event.pageX
    }

    function stopExpand() {
        expanding = false
        localStorage.setItem(STORE_KEY, width.toString())
    }

    function expand(event: MouseEvent) {
        if (!expanding) return
        width = event.pageX
        dispatchResize()
    }

    onMount(dispatchResize)
</script>

<svelte:window on:mouseup={stopExpand} on:mousemove={expand} />

<div class="flex flex-row h-full">
    <div
        class=" bg-gray-50  flex flex-col gap-1 h-full overflow-y-scroll py-1"
        style="min-width: 350px; width: {width}px"
        data-cy="explorer"
        on:click={onClick}
        transition:fly|local={{ x: -400, duration: 200, opacity: 1 }}
    >
        <slot />
        <span class="py-10" />
    </div>

    <div
        data-cy="explorer-expander"
        class=" cursor-col-resize h-full border-gray-200 border"
        on:mousedown={startExpand}
        on:mousemove={expand}
    />
</div>
