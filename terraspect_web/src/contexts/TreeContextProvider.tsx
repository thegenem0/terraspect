import { createContext, useContext, useEffect, useState } from 'react'

import { DataNode, useGraphQuery } from '@/hooks/useGraphQuery'

export type TreeContext = {
  treeData: DataNode[]
  isLoading: boolean
  activeNode: DataNode | undefined
  toggleActiveNodeById: (id: string) => void
}

export const TreeContext = createContext<TreeContext>({} as TreeContext)

export const useTreeContext = () => useContext(TreeContext)

type Props = {
  children: React.ReactNode
}

export const TreeContextProvider = ({ children }: Props) => {
  const [treeData, setTreeData] = useState<DataNode[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [activeNode, setActiveNode] = useState<DataNode | undefined>()

  const {
    data,
    isLoading: queryLoading,
    isFetching: queryFetching,
    isError
  } = useGraphQuery()

  useEffect(() => {
    if (queryLoading || queryFetching) {
      setIsLoading(true)
    } else {
      setIsLoading(false)
    }

    if ((!queryLoading && !queryFetching && isError) || !data) {
      setTreeData([])
    }

    if (data) {
      setTreeData(data.tree)
    }
  }, [data, queryLoading, queryFetching, isError])

  const toggleActiveNodeById = (id: string) => {
    const findNodeById = (
      nodes: DataNode[],
      id: string
    ): DataNode | undefined => {
      for (const node of nodes) {
        if (node.id === id) {
          return node
        }
        if (node.children) {
          const found = findNodeById(node.children, id)
          if (found) {
            return found
          }
        }
      }
      return undefined
    }

    const node = findNodeById(treeData, id)
    if (node) {
      if (activeNode?.id === node.id) {
        setActiveNode(undefined)
      } else {
        setActiveNode(node)
      }
    }
  }

  const treeContext = {
    treeData,
    isLoading,
    activeNode,
    toggleActiveNodeById
  }

  return (
    <TreeContext.Provider value={treeContext}>{children}</TreeContext.Provider>
  )
}
