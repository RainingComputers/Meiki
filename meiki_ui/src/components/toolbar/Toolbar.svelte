<script lang="ts">
    import { fly } from "svelte/transition"
</script>

<nav
    class="flex flex-row w-full h-12 drop-shadow-card bg-slate-800 items-center px-4 py-7 gap-2"
    transition:fly={{ y: -50, duration: 350, opacity: 1 }}
>
    <slot />
</nav>
