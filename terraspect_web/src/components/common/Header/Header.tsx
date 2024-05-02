import { UserButton } from '@clerk/clerk-react'
import { Link } from '@tanstack/react-router'
import { ReactNode } from 'react'

import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger
} from '@/components/ui/navigation-menu'
import { cn } from '@/lib/utils'

const Header = () => {
  return (
    <div className="bg-slate-800">
      <div className="mx-auto flex flex-row justify-between gap-4 px-12 py-2">
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              <NavigationMenuTrigger>Terraform</NavigationMenuTrigger>
              <NavigationMenuContent>
                <ul className="grid gap-3 p-6 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]">
                  <Link to="/graph" className="no-underline">
                    <ListItem title="Graph View">
                      Visualize your infrastructure.
                    </ListItem>
                  </Link>
                </ul>
              </NavigationMenuContent>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <NavigationMenuTrigger>Account</NavigationMenuTrigger>
              <NavigationMenuContent>
                <ul className="grid gap-3 p-6 md:w-[400px] lg:w-[500px] lg:grid-cols-[.75fr_1fr]">
                  <Link to="/keys" className="no-underline">
                    <ListItem title="API Keys">Manage your API keys.</ListItem>
                  </Link>
                </ul>
              </NavigationMenuContent>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>

        <UserButton />
      </div>
    </div>
  )
}

export default Header

type ListItemProps = {
  className?: string
  title: string
  children: ReactNode
}

const ListItem = ({ className, title, children }: ListItemProps) => {
  return (
    <li>
      <NavigationMenuLink asChild>
        <div
          className={cn(
            'block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground',
            className
          )}
        >
          <div className="text-sm font-medium leading-none">{title}</div>
          <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
            {children}
          </p>
        </div>
      </NavigationMenuLink>
    </li>
  )
}
