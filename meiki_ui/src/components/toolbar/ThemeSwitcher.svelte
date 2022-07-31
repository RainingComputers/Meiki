<script lang="ts">
    import { onMount } from "svelte"
    import currentTheme from "$lib/stores/theme"
    import SunIcon from "$cmp/icons/SunIcon.svelte"
    import MoonIcon from "$cmp/icons/MoonIcon.svelte"

    let isDarkTheme: boolean = false

    function setTheme(isDarkTheme: boolean) {
        if (isDarkTheme) {
            document.documentElement.classList.add("dark")
            localStorage.theme = "dark"
            $currentTheme = "dark"
        } else {
            document.documentElement.classList.remove("dark")
            localStorage.theme = "light"
            $currentTheme = "light"
        }
    }

    function isCurrentThemeDark() {
        const prefersDarkTheme = window.matchMedia("(prefers-color-scheme: dark)").matches
        const noThemeSet = !("theme" in localStorage)

        return localStorage.theme === "dark" || (noThemeSet && prefersDarkTheme)
    }

    function flipTheme() {
        isDarkTheme = !isDarkTheme
        setTheme(isDarkTheme)
    }

    onMount(() => {
        isDarkTheme = isCurrentThemeDark()
        setTheme(isDarkTheme)
    })
</script>

<div
    class="flex gap-2 flex-row items-center hover:bg-toolbar-focus rounded-xl p-2 cursor-pointer"
    on:click={flipTheme}
    data-cy="themeSwitcher"
>
    <span class="stroke-toolbar-content h-6 w-6">
        {#if isDarkTheme}
            <SunIcon />
        {:else}
            <MoonIcon />
        {/if}
    </span>
</div>
