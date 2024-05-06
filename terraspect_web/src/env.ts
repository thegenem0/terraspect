import { createEnv } from '@t3-oss/env-core'
import { z } from 'zod'

export const env = createEnv({
  clientPrefix: 'VITE_',
  client: {
    VITE_CLERK_PUBLISHABLE_KEY: z.string(),
    VITE_API_BASE_URI: z.string()
  },
  runtimeEnv: import.meta.env
})
