import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'

import DetailsDialog from '@/components/pages/Graph/DetailsDialog/DetailsDialog'
import GraphContainer from '@/components/pages/Graph/NodeTree/NodeTree'
import {
  TreeContextProvider,
  useTreeContext
} from '@/contexts/TreeContextProvider'

const GraphSearchValidator = z.object({
  projectId: z.string().catch(''),
  planId: z.string().catch('')
})

export const Route = createFileRoute('/_layout/graph')({
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
      <div className="mx-12 flex min-h-[85vh] flex-row gap-4">
        <div className="w-full overflow-hidden rounded-xl bg-white ">
          <GraphRenderer />
        </div>
        <DetailsDialog />
      </div>
    </TreeContextProvider>
  )
}

const GraphRenderer = () => {
  const { isLoading } = useTreeContext()
  return isLoading ? <GraphLoaderSkeleton /> : <GraphContainer />
}

const GraphLoaderSkeleton = () => {
  return (
    <div className="flex size-full items-center justify-center bg-gray-100 p-4">
      <svg
        width="100%"
        height="100%"
        viewBox="0 0 400 200"
        className="animate-pulse"
      >
        <circle cx="50" cy="50" r="10" fill="#cbd5e1" />
        <circle cx="150" cy="150" r="10" fill="#cbd5e1" />
        <circle cx="250" cy="50" r="10" fill="#cbd5e1" />
        <circle cx="350" cy="150" r="10" fill="#cbd5e1" />

        <line
          x1="50"
          y1="50"
          x2="150"
          y2="150"
          stroke="#cbd5e1"
          strokeWidth="2"
        />
        <line
          x1="150"
          y1="150"
          x2="250"
          y2="50"
          stroke="#cbd5e1"
          strokeWidth="2"
        />
        <line
          x1="250"
          y1="50"
          x2="350"
          y2="150"
          stroke="#cbd5e1"
          strokeWidth="2"
        />
      </svg>
    </div>
  )
}
