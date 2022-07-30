/// <reference types="@sveltejs/kit" />
/// <reference types="vite/client" />

declare module "brace/mode/markdown"
declare module "brace/theme/textmate"
declare module "brace/theme/monokai"
declare module "brace/theme/tomorrow_night_bright"

interface ImportMetaEnv {
    readonly VITE_MEIKI_SERVER_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
