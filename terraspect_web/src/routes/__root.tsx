import { ClerkProvider } from '@clerk/clerk-react'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { createRootRoute, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

import { env } from '@/env'

export const Route = createRootRoute({
  component: () => <DefaultLayout />
})

const DefaultLayout = () => {
  return (
    <>
      <ClerkProvider publishableKey={env.VITE_CLERK_PUBLISHABLE_KEY}>
        <Outlet />
      </ClerkProvider>
      <TanStackRouterDevtools />
      <ReactQueryDevtools />
    </>
  )
}
