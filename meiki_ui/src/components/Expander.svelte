<script lang="ts">
    export let name: string
    export let attachedElement: HTMLDivElement

    const STORE_KEY = name + "-width"

    export let widthPercentage: number = 25

    function getInitialWidth() {
        const defaultWidth = (widthPercentage * window.innerWidth) / 100
        const storedWidth = parseFloat(localStorage.getItem(STORE_KEY))
        return storedWidth || defaultWidth
    }

    let width: number = getInitialWidth()
    let expanding = false

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
        attachedElement.style.width = width + "px"
    }
</script>

<svelte:window on:mouseup={stopExpand} on:mousemove={expand} />

<div
    data-cy={name}
    class="cursor-col-resize h-full bg-base-2 w-px"
    on:mousedown={startExpand}
    on:mousemove={expand}
/>
