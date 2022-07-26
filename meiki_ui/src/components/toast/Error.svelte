<script lang="ts">
    import { scale } from "svelte/transition"

    export let fullWidth = false
</script>

<span
    class=" text-content-error bg-background-error p-2 rounded-lg border-2 border-border-error drop-shadow-card"
    class:w-full={fullWidth}
    transition:scale={{ duration: 70, opacity: 1 }}
>
    <slot />
</span>
