import {
  SigmaContainer,
  useLoadGraph,
  useRegisterEvents
} from '@react-sigma/core'
import Graph from 'graphology'
import { useEffect } from 'react'

import { useTreeContext } from '@/contexts/TreeContextProvider'
import { DataNode, Variable } from '@/hooks/useGraphQuery'

type GraphNode = {
  id: string
  x: number
  y: number
  size: number
  label: string
  children?: GraphNode[]
  variables?: Variable[]
}

interface LoadGraphProps {
  data: DataNode[]
}

export const LoadGraph = ({ data }: LoadGraphProps) => {
  const loadGraph = useLoadGraph()
  const registerEvents = useRegisterEvents()
  const { toggleActiveNodeById, activeNode } = useTreeContext()

  useEffect(() => {
    const graph = new Graph<GraphNode>()
    registerEvents({
      clickNode: (event) => toggleActiveNodeById(event.node)
    })

    const addNodesRecursively = (
      node: DataNode,
      graph: Graph,
      level = 0,
      parentNodeId?: string,
      siblingIndex = 0,
      totalSiblings = 1
    ) => {
      const spacingX = 50
      const spacingY = 100

      const posX =
        siblingIndex * spacingX - ((totalSiblings - 1) * spacingX) / 2

      const posY = (4 - level) * spacingY

      graph.addNode(node.id, {
        id: node.id,
        x: posX,
        y: posY,
        size: 40 / (level + 1),
        label: node.label,
        variables: node.variables
      })

      if (parentNodeId) {
        graph.addEdge(parentNodeId, node.id)
      }

      if (node.children) {
        const numChildren = node.children.length
        node.children.forEach((child, index) => {
          addNodesRecursively(
            child,
            graph,
            level + 1,
            node.id,
            index,
            numChildren
          )
        })
      }
    }
    for (const node of data) {
      const graphNode = {
        id: node.id,
        x: Math.random() * 10000,
        y: Math.random() * 10000,
        size: 30,
        label: node.label,
        variables: node.variables,
        children: node.children
      }
      addNodesRecursively(graphNode, graph)
    }

    loadGraph(graph)
  }, [loadGraph, registerEvents, data, toggleActiveNodeById, activeNode])

  // Returning null to get a valid component
  return null
}

export const DisplayGraph = () => {
  const { treeData, isLoading } = useTreeContext()

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <SigmaContainer
      settings={{ allowInvalidContainer: true }}
      style={{ height: '100%' }}
    >
      <LoadGraph data={treeData} />
    </SigmaContainer>
  )
}
