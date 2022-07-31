const colors = require("tailwindcss/colors")

module.exports = {
    darkMode: "class",

    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        // prettier-ignore
        colors: {
            "transparent": "#00000000",
            "overlay": "#4b5563",

            "base-0": "#ffffff",
            "base-1": "#f9fafb",
            "base-2": "#e5e7eb",
            "base-3": "#d1d5db",
            "base-4": "#9ca3af",
            "content": "#0f172a",

            "toolbar": "#1f2937",
            "toolbar-focus": "#475569",
            "toolbar-content": "#e5e7eb",
            
            "primary": "#2563eb",
            "primary-focus": "#3b82f6",
            "primary-content": "#ffffff",
            
            "error": "#fef2f2",
            "error-outline": "#f87171",
            "error-content": "#ef4444",

            "success": "#f0fdf4",
            "success-outline": "#4ade80",
            "success-content": "#22c55e",
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
