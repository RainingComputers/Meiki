<script lang="ts">
    import { scale } from "svelte/transition"

    export let fullWidth = false
</script>

<span
    class=" text-red-500 bg-red-50 p-2 rounded-lg border-2 border-red-400 drop-shadow-card"
    class:w-full={fullWidth}
    transition:scale={{ duration: 70, opacity: 1 }}
>
    <slot />
</span>
