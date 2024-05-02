import { useQueryClient } from '@tanstack/react-query'
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState
} from 'react'

import { DataNode, useGraphQuery } from '@/hooks/queries/useGraphQuery'

export type TreeContext = {
  treeData: DataNode[]
  activeNode: DataNode | undefined
  hoveredNodeId: string | undefined
  toggleActiveNodeById: (id: string) => void
  setHoveredNodeId: (id?: string) => void
  refreshTree: () => void
}

export const TreeContext = createContext<TreeContext>({} as TreeContext)

export const useTreeContext = () => useContext(TreeContext)

type Props = {
  projectId?: string
  planId?: string
  children: React.ReactNode
}

export const TreeContextProvider = ({ projectId, planId, children }: Props) => {
  const [activeNode, setActiveNode] = useState<DataNode | undefined>()
  const [hoveredNodeId, setHoveredNodeId] = useState<string | undefined>()

  const queryClient = useQueryClient()

  const {
    data,
    isLoading: queryLoading,
    isFetching: queryFetching,
    refetch: refetchTree
  } = useGraphQuery({
    projectId,
    planId
  })

  const treeData = useMemo(() => {
    if (!data?.tree?.nodes) {
      const emptyNode = {
        tree: {
          nodes: []
        }
      }
      return emptyNode.tree.nodes
    } else {
      return data?.tree.nodes
    }
  }, [data])

  useEffect(() => {
    if (projectId && planId) {
      refetchTree()
    }
  }, [projectId, planId, refetchTree])

  const findNodeById = useCallback(
    (id: string, nodes?: DataNode[]): DataNode | undefined => {
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
        prev?.id === id ? undefined : findNodeById(id, data?.tree.nodes)
      )
    },
    [findNodeById, data]
  )
  const refreshTree = useCallback(() => {
    queryClient.invalidateQueries({ queryKey: ['tree'] })
  }, [queryClient])

  const treeContext = {
    treeData,
    activeNode,
    hoveredNodeId,
    toggleActiveNodeById,
    setHoveredNodeId,
    refreshTree
  }

  if (queryLoading || queryFetching) {
    return <div>Loading...</div>
  }

  return (
    <TreeContext.Provider value={treeContext}>{children}</TreeContext.Provider>
  )
}
