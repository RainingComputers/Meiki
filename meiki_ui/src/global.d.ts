/// <reference types="@sveltejs/kit" />
/// <reference types="vite/client" />

declare module "brace/mode/markdown"
declare module "brace/theme/textmate"
declare module "brace/theme/monokai"
declare module "brace/theme/tomorrow_night_bright"
declare module "brace/theme/twilight"

interface ImportMetaEnv {
    readonly VITE_MEIKI_SERVER_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
