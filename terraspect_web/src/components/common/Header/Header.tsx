import { UserButton } from '@clerk/clerk-react'
import { Link } from '@tanstack/react-router'

const Header = () => {
  return (
    <div className="bg-slate-800">
      <div className="mx-auto flex flex-row justify-evenly gap-4 px-12 py-2">
        <img src="/terraspect-logo.png" alt="logo" className="h-12" />
        <nav className="hidden items-center gap-6 text-lg font-semibold text-white md:flex">
          <Link className="hover:underline" to="/projects">
            Projects
          </Link>
          <Link className="hover:underline" to="/keys">
            API Keys
          </Link>
        </nav>
        <UserButton />
      </div>
    </div>
  )
}

export default Header
