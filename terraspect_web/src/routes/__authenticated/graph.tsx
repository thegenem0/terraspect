import { createFileRoute } from '@tanstack/react-router'

import { DisplayGraph } from '@/components/NodeTree/NodeTree'
import DetailsDialog from '@/components/pages/Graph/DetailsDialog/DetailsDialog'
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
