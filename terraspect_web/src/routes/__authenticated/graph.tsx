import { createFileRoute } from '@tanstack/react-router'

import DetailsDialog from '@/components/pages/Graph/DetailsDialog/DetailsDialog'
import { DisplayGraph } from '@/components/pages/Graph/NodeTree/NodeTree'
import { TreeContextProvider } from '@/contexts/TreeContextProvider'

export const Route = createFileRoute('/__authenticated/graph')({
  component: () => <GraphComponent />
})

const GraphComponent = () => {
  return (
    <TreeContextProvider>
      <div className="flex h-full flex-row gap-4">
        <div className="w-full overflow-hidden">
          <DisplayGraph />
        </div>
        <DetailsDialog />
      </div>
    </TreeContextProvider>
  )
}
