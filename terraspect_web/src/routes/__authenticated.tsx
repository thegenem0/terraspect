import { SignedIn, SignedOut, SignIn } from '@clerk/clerk-react'
import { createFileRoute, Outlet } from '@tanstack/react-router'

import Header from '@/components/common/Header/Header'

export const Route = createFileRoute('/__authenticated')({
  component: () => <AuthenticatedLayout />
})

const AuthenticatedLayout = () => {
  return (
    <div className="flex size-full flex-col">
      <SignedOut>
        <SignIn />
      </SignedOut>
      <SignedIn>
        <Header />
        <Outlet />
      </SignedIn>
    </div>
  )
}
