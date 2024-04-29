import { UserButton } from '@clerk/clerk-react'

const Header = () => {
  return (
    <nav className="flex w-full flex-row justify-between gap-4 bg-sky-500 px-12 py-4">
      <h1>Header</h1>
      <UserButton />
    </nav>
  )
}

export default Header
