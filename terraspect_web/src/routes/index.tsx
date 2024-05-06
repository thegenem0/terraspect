import { createFileRoute, redirect } from '@tanstack/react-router'

// Default root redirect for triggering auth flow
export const Route = createFileRoute('/')({
  loader: () => {
    throw redirect({
      to: '/projects'
    })
  }
})
