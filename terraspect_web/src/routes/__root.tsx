import { ClerkProvider, SignedIn, SignedOut, SignIn } from '@clerk/clerk-react'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { createRootRoute, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { icons } from 'lucide-react'
import { Toaster } from 'react-hot-toast'

import { env } from '@/env'

const toastOptions = {
  duration: 4000,
  success: {
    style: {
      background: '#61c0ad',
      color: '#11012B'
    },
    icons: {
      success: icons.Check
    }
  },
  error: {
    style: {
      background: '#EF887E',
      color: '#11012B'
    },
    icons: {
      error: icons.X
    }
  }
}

export const Route = createRootRoute({
  component: () => <DefaultLayout />
})

// Authentication handling with Clerk Provider
const DefaultLayout = () => {
  return (
    <div className="h-full bg-default bg-cover">
      <Toaster position="top-center" toastOptions={toastOptions} />
      <ClerkProvider publishableKey={env.VITE_CLERK_PUBLISHABLE_KEY}>
        <SignedIn>
          <Outlet />
        </SignedIn>
        <SignedOut>
          <div className="flex h-full flex-row items-center justify-center">
            <SignIn routing="virtual" />
          </div>
        </SignedOut>
      </ClerkProvider>
      {import.meta.env.DEV && (
        <>
          <TanStackRouterDevtools />
          <ReactQueryDevtools />
        </>
      )}
    </div>
  )
}
