<script lang="ts">
    import { fly } from "svelte/transition"

    export let width: string = ""
    export let expand: boolean = true
</script>

<nav
    class="flex flex-row h-12 drop-shadow-card bg-slate-800 items-center px-4 py-7 gap-2"
    class:flex-grow={expand}
    style="min-width: 350px; width: {width};"
    in:fly={{ y: -50, duration: 200, opacity: 1 }}
>
    <slot />
</nav>
