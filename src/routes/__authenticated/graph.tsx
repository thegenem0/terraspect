import { createFileRoute } from '@tanstack/react-router'

import DetailsPanel from '@/components/DetailsPanel/DetailsPanel'
import { DisplayGraph } from '@/components/NodeTree/NodeTree'
import { TreeContextProvider } from '@/contexts/TreeContextProvider'

export const Route = createFileRoute('/__authenticated/graph')({
  component: () => <GraphComponent />
})

const GraphComponent = () => {
  return (
    <TreeContextProvider>
      <div className="flex h-full flex-row gap-4">
        <div className="w-9/12 overflow-hidden">
          <DisplayGraph />
        </div>
        <div className="w-3/12 overflow-auto">
          <DetailsPanel />
        </div>
      </div>
    </TreeContextProvider>
  )
}

