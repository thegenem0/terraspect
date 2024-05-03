import { MultiDirectedGraph } from 'graphology'

import { TreeDataNode } from '@/contexts/TreeContextProvider'

export type GraphNode = {
  id: string
  x: number
  y: number
  label: string
  size: number
  color: string
  highlighted?: boolean
}

export type GraphEdge = {
  label: string
}

interface BuildGraphProps {
  data: TreeDataNode[]
}

export const buildGraph = ({ data }: BuildGraphProps) => {
  const graph = new MultiDirectedGraph<GraphNode, GraphEdge>()

  for (const node of data) {
    addNodesRecursively(node, graph)
  }

  return graph
}

const addNodesRecursively = (
  node: TreeDataNode,
  graph: MultiDirectedGraph<GraphNode, GraphEdge>,
  level = 0,
  parentNodeId?: string
) => {
  graph.addNode(node.id, {
    id: node.id,
    x: Math.random(),
    y: Math.random(),
    size: calculateNodeSize(level),
    label: node.label,
    color: getNodeColor(level),
    highlighted: false
  })

  if (parentNodeId) {
    graph.addDirectedEdge(parentNodeId, node.id)
  }

  if (node.children) {
    node.children.forEach((child) => {
      addNodesRecursively(child, graph, level + 1, node.id)
    })
  }
}

const calculateNodeSize = (level: number) => {
  const baseSize = 50
  if (level === 0) return baseSize / 2
  if (level === 1) return baseSize / 3
  return baseSize / (level * level)
}

const getNodeColor = (level: number) => {
  if (level === 0) return 'red'
  if (level === 1) return 'blue'
  if (level === 2) return 'green'
  return 'grey'
}
