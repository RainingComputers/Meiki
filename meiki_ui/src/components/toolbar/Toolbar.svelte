<script lang="ts">
    import { fly } from "svelte/transition"
</script>

<nav
    class="flex flex-row w-full h-12 drop-shadow-card bg-background-toolbar items-center px-4 py-7 gap-2"
    in:fly={{ y: -50, duration: 200, opacity: 1 }}
>
    <slot />
</nav>
