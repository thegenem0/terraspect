/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_CLERK_PUBLISHABLE_KEY: string
  readonly VITE_API_BASE_URI: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
