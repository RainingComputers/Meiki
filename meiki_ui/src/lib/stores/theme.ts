import { writable } from "svelte/store"

export const theme = writable("dark")

export function setTheme(dark: boolean) {
    if (dark) {
        document.documentElement.setAttribute("data-theme", "dark")
        localStorage.theme = "dark"
        theme.set("dark")
    } else {
        document.documentElement.setAttribute("data-theme", "light")
        localStorage.theme = "light"
        theme.set("light")
    }
}

export function isCurrentThemeDark() {
    const prefersDarkTheme = window.matchMedia("(prefers-color-scheme: dark)").matches
    const noThemeSet = !("theme" in localStorage)

    return localStorage.theme === "dark" || (noThemeSet && prefersDarkTheme)
}

export function flipTheme() {
    setTheme(localStorage.theme != "dark")
}

export function initTheme() {
    setTheme(isCurrentThemeDark())
}
