<script lang="ts">
    import { scale } from "svelte/transition"

    export let width: number = 400
</script>

<span
    class=" bg-gray-50 border-gray-200 border drop-shadow-card "
    style="width: {width}px;"
    transition:scale={{ duration: 30, opacity: 1 }}
>
    <slot />
</span>
