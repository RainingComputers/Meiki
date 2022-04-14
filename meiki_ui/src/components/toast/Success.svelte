<script lang="ts">
    import { scale } from "svelte/transition"

    export let fullWidth = false
</script>

<span
    class=" text-green-500 bg-green-50 p-4 rounded-lg border-2 border-green-400 drop-shadow-card"
    class:w-full={fullWidth}
    transition:scale={{ duration: 70, opacity: 1 }}
>
    <slot />
</span>
