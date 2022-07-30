<script lang="ts">
    import { onMount } from "svelte"
    import currentTheme from "$lib/stores/theme"
    import SunIcon from "$cmp/icons/SunIcon.svelte"
    import MoonIcon from "$cmp/icons/MoonIcon.svelte"

    export let isDarkTheme: boolean = false

    export function setTheme(isDarkTheme: boolean) {
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
        return (
            localStorage.theme === "dark" ||
            (!("theme" in localStorage) &&
                window.matchMedia("(prefers-color-scheme: dark)").matches)
        )
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
    class="flex gap-2 flex-row items-center hover:bg-toolbarFocus rounded-xl p-2 cursor-pointer"
    on:click={flipTheme}
    data-cy="themeSwitcher"
>
    <span class="stroke-contentTitle h-6 w-6">
        {#if isDarkTheme}
            <SunIcon />
        {:else}
            <MoonIcon />
        {/if}
    </span>
</div>
