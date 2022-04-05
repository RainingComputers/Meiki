/// <reference types="@sveltejs/kit" />
/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_MEIKI_SERVER_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
