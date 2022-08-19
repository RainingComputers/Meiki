import { sveltekit } from "@sveltejs/kit/vite"
import { resolve } from "path"

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit()],
    resolve: {
        alias: {
            // these are the aliases and paths to them
            $cmp: resolve("src/components"),
            // $lib: resolve("src/lib"),
            $data: resolve("src/data"),
            $root: resolve(""),
            // '@utils': path.resolve('./src/lib/utils')
        },
    },
}

export default config
