const colors = require("tailwindcss/colors")

module.exports = {
    darkMode: "media",

    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        color: {
            white: colors.white,
            gray: colors.gray,
            yellow: colors.yellow,
            red: colors.red,
            blue: colors.blue,
            black: colors.black,
        },
        fontFamily: {
            serif: ["Montserrat", "system-ui"],
        },
        extend: {
            dropShadow: {
                card: "0 0px 2px rgb(0 0 0 / 0.1)",
            },
        },
    },
    variants: {},
    plugins: [],
}
