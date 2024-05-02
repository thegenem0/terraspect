import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'

import DetailsDialog from '@/components/pages/Graph/DetailsDialog/DetailsDialog'
import GraphContainer from '@/components/pages/Graph/NodeTree/NodeTree'
import { TreeContextProvider } from '@/contexts/TreeContextProvider'

const GraphSearchValidator = z.object({
  projectId: z.string().catch(''),
  planId: z.string().catch('')
})

export const Route = createFileRoute('/__authenticated/graph')({
  validateSearch: GraphSearchValidator,
  component: () => <GraphComponent />
})

const GraphComponent = () => {
  const { projectId, planId } = Route.useSearch()

  return (
    <TreeContextProvider
      key={`${projectId}-${planId}`}
      projectId={projectId}
      planId={planId}
    >
      <div className="flex h-full flex-row gap-4">
        <div className="w-full overflow-hidden">
          <GraphContainer />
        </div>
        <DetailsDialog />
      </div>
    </TreeContextProvider>
  )
}
