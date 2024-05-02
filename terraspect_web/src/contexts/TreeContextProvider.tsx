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
  isLoading: boolean
  activeNode: DataNode | undefined
  hoveredNodeId: string | undefined
  toggleActiveNodeById: (id: string) => void
  setHoveredNodeId: (id?: string) => void
  refreshTree: () => void
}

export const TreeContext = createContext<TreeContext>({} as TreeContext)

export const useTreeContext = () => useContext(TreeContext)

type Props = {
  children: React.ReactNode
}

export const TreeContextProvider = ({ children }: Props) => {
  const [isLoading, setIsLoading] = useState(true)
  const [activeNode, setActiveNode] = useState<DataNode | undefined>()
  const [hoveredNodeId, setHoveredNodeId] = useState<string | undefined>()

  const queryClient = useQueryClient()

  const {
    data,
    isLoading: queryLoading,
    isFetching: queryFetching
  } = useGraphQuery()

  useEffect(() => {
    if (queryLoading || queryFetching) {
      setIsLoading(true)
    } else {
      setIsLoading(false)
    }
  }, [queryLoading, queryFetching])

  const treeData = useMemo(() => {
    if (!data) return []
    return data.tree.nodes
  }, [data])

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
    isLoading,
    activeNode,
    hoveredNodeId,
    toggleActiveNodeById,
    setHoveredNodeId,
    refreshTree
  }

  return (
    <TreeContext.Provider value={treeContext}>{children}</TreeContext.Provider>
  )
}
