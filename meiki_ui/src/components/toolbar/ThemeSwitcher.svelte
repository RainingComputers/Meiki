<script lang="ts">
    import { onMount } from "svelte"
    import { theme, initTheme, flipTheme } from "$lib/stores/theme"
    import SunIcon from "$cmp/icons/SunIcon.svelte"
    import MoonIcon from "$cmp/icons/MoonIcon.svelte"

    onMount(initTheme)
</script>

<div
    class="flex gap-2 flex-row items-center hover:bg-toolbar-focus rounded-xl p-2 cursor-pointer"
    on:click={flipTheme}
    data-cy="themeSwitcher"
>
    <span class="stroke-toolbar-content h-6 w-6">
        {#if $theme == "dark"}
            <SunIcon />
        {:else}
            <MoonIcon />
        {/if}
    </span>
</div>
