const colors = require("tailwindcss/colors")

module.exports = {
    darkMode: "media",

    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        colors: {
            "transparent": "#00000000",
            "searchbar": "#ffffff",
            "info": "#0f172a",
            "error": "#fef2f2",
            "overlay": "#4b5563",
            "panel": "#f9fafb",
            "primaryFocus": "#3b82f6",
            "primary": "#2563eb",
            "secondaryFocus": "#e5e7eb",
            "secondary": "#d1d5db",
            "success": "#f0fdf4",
            "toolbarFocus": "#475569",
            "toolbar": "#1f2937",
            "borderError": "#f87171",
            "borderPanel": "#e5e7eb",
            "borderSuccess": "#4ade80",
            "contentErrorWatermark": "#7f1d1d",
            "contentError": "#ef4444",
            "contentPlaceholder": "#9ca3af",
            "contentPrimary": "#ffffff",
            "contentSecondary": "334155",
            "contentSuccess": "#22c55e",
            "contentTitle": "#e5e7eb",
        },
        fontFamily: {
            serif: ["Montserrat", "system-ui"],
        },
        extend: {
            dropShadow: {
                card: "0 0px 2px rgb(0 0 0 / 0.1)",
            },
            opacity: {
                4: ".04",
            },
        },
    },
    variants: {},
    plugins: [require("@tailwindcss/typography")],
}
