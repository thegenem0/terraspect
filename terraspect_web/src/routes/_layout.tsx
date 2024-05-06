import { createFileRoute, Outlet } from '@tanstack/react-router'

import Header from '@/components/common/Header/Header'

export const Route = createFileRoute('/_layout')({
  component: AuthenticatedLayout
})

// Apply layout to authenticated routes
function AuthenticatedLayout() {
  return (
    <>
      <Header />
      <div className="py-6">
        <Outlet />
      </div>
    </>
  )
}
