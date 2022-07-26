<script lang="ts">
    import { scale } from "svelte/transition"

    export let width: number = 400
</script>

<span
    class=" bg-background-panel border-border-panel border drop-shadow-card "
    style="width: {width}px;"
    transition:scale={{ duration: 70, opacity: 1 }}
>
    <slot />
</span>
