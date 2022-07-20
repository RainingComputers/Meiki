/// <reference types="@sveltejs/kit" />
/// <reference types="vite/client" />

declare module "brace/mode/markdown"
declare module "brace/theme/textmate"

interface ImportMetaEnv {
    readonly VITE_MEIKI_SERVER_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
