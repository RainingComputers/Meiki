const colors = require("tailwindcss/colors")

module.exports = {
    darkMode: "media",

    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        colors: {
            "type-semantics-modifier": "#000000", // testing
            white: "#ffffff",
            info: "#0f172a",
            "background-error": "#fef2f2",
            "background-overlay": "#4b5563",
            "background-panel": "#f9fafb",
            "background-primary-focus": "#3b82f6",
            "background-primary": "#2563eb",
            "background-secondary-focus": "#e5e7eb",
            "background-secondary": "e2e8f0",
            "background-success": "#f0fdf4",
            "background-toolbar-focus": "#475569",
            "background-toolbar": "#1f2937",
            "border-error": "#f87171",
            "border-panel": "#e5e7eb",
            "border-success": "#4ade80",
            "content-error-watermark": "#7f1d1d",
            "content-error": "#ef4444",
            "content-placeholder": "#9ca3af",
            "content-primary": "#ffffff",
            "content-secondary": "334155",
            "content-success": "#22c55e",
            "content-title": "#e5e7eb",
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
