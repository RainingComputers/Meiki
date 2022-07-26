<script lang="ts">
    import { scale } from "svelte/transition"

    export let fullWidth = false
</script>

<span
    class=" text-content-success bg-background-success p-4 rounded-lg border-2 border-border-success drop-shadow-card"
    class:w-full={fullWidth}
    transition:scale={{ duration: 70, opacity: 1 }}
>
    <slot />
</span>
