/// <reference types="@sveltejs/kit" />
/// <reference types="vite/client" />
/// <reference types="brace/mode/markdown" />
/// <reference types="brace/theme/textmate" />

interface ImportMetaEnv {
    readonly VITE_MEIKI_SERVER_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
