import { SignedIn, SignedOut, SignIn } from '@clerk/clerk-react'
import { createFileRoute, Outlet } from '@tanstack/react-router'

import Header from '@/components/Header/Header'

export const Route = createFileRoute('/__authenticated')({
  component: () => <AuthenticatedLayout />
})

const AuthenticatedLayout = () => {
  return (
    <>
      <SignedOut>
        <SignIn />
      </SignedOut>
      <SignedIn>
        <Header />
        <Outlet />
      </SignedIn>
    </>
  )
}
