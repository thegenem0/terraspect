import { useQueryClient } from '@tanstack/react-query'
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState
} from 'react'

import { ChangeItem, useChangesQuery } from '@/hooks/queries/useChangesQuerry'
import { DataNode, useGraphQuery } from '@/hooks/queries/useGraphQuery'
import { mergeChangesForModule } from '@/lib/mergeChangesForModule'

export type TreeContext = {
  treeData: TreeDataNode[]
  isLoading?: boolean
  activeNode: TreeDataNode | undefined
  hoveredNodeId: string | undefined
  toggleActiveNodeById: (id: string) => void
  setHoveredNodeId: (id?: string) => void
  refreshTree: () => void
}

export const TreeContext = createContext<TreeContext>({} as TreeContext)

export const useTreeContext = () => useContext(TreeContext)

export type TreeDataNode = DataNode & {
  changes?: ChangeItem[]
}

type Props = {
  projectId?: string
  planId?: string
  children: React.ReactNode
}

export const TreeContextProvider = ({ projectId, planId, children }: Props) => {
  const [activeNode, setActiveNode] = useState<TreeDataNode | undefined>()
  const [hoveredNodeId, setHoveredNodeId] = useState<string | undefined>()
  const [isLoading, setIsLoading] = useState(true)

  const queryClient = useQueryClient()

  const {
    data: graphDataResult,
    isLoading: queryLoading,
    isFetching: queryFetching,
    refetch: refetchTree
  } = useGraphQuery({
    projectId,
    planId
  })

  const {
    data: changesDataResult,
    isLoading: changesQueryLoading,
    isFetching: changesQueryFetching,
    refetch: refetchChanges
  } = useChangesQuery({
    projectId,
    planId
  })

  const treeDataNodes = useMemo(() => {
    const treeData = graphDataResult?.tree?.nodes || []
    const changesData = changesDataResult?.changes || []

    return mergeChangesForModule(treeData, changesData)
  }, [graphDataResult, changesDataResult])

  useEffect(() => {
    if (projectId && planId) {
      refetchTree()
      refetchChanges()
    }
  }, [projectId, planId, refetchTree, refetchChanges])

  const findNodeById = useCallback(
    (id: string, nodes?: TreeDataNode[]): TreeDataNode | undefined => {
      if (!nodes) return undefined
      for (const node of nodes) {
        if (node.id === id) return node
        if (node.children) {
          const found = findNodeById(id, node.children)
          if (found) return found
        }
      }
      return undefined
    },
    []
  )

  const toggleActiveNodeById = useCallback(
    (id: string) => {
      setActiveNode((prev) =>
        prev?.id === id ? undefined : findNodeById(id, treeDataNodes)
      )
    },
    [findNodeById, treeDataNodes]
  )
  const refreshTree = useCallback(() => {
    queryClient.invalidateQueries({ queryKey: ['tree', 'changes'] })
  }, [queryClient])

  const treeContext = {
    treeData: treeDataNodes,
    isLoading:
      queryLoading ||
      queryFetching ||
      changesQueryLoading ||
      changesQueryFetching,
    activeNode,
    hoveredNodeId,
    toggleActiveNodeById,
    setHoveredNodeId,
    refreshTree
  }

  // if (
  //   queryLoading ||
  //   queryFetching ||
  //   changesQueryLoading ||
  //   changesQueryFetching
  // ) {
  //   return <div>Loading...</div>
  // }

  return (
    <TreeContext.Provider value={treeContext}>{children}</TreeContext.Provider>
  )
}
