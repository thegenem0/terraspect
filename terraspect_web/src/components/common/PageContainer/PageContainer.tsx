import { ReactNode } from 'react'

import { cn } from '@/lib/utils'

type PageContainerProps = {
  children: ReactNode | ReactNode[]
  className?: string
}

const PageContainer = ({ children, className }: PageContainerProps) => {
  return (
    <div className="flex h-full flex-row justify-center py-2">
      <div className="flex w-2/3 flex-col items-center gap-8 rounded-lg bg-slate-600 py-8">
        <div className={cn('flex flex-col justify-center', className)}>
          {children}
        </div>
      </div>
    </div>
  )
}

export default PageContainer
