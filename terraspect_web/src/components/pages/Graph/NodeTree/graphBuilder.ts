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

const nodeColors = {
  purple: '#7B42BC',
  red: '#97454d',
  blue: '#132e65',
  green: '#168581',
  gray: '#757687'
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
    color: getNodeColor(level, node.changes && node.changes.length > 0),
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

const getNodeColor = (level: number, hasChange?: boolean) => {
  if (hasChange) return nodeColors.red
  if (level === 0) return nodeColors.purple
  if (level === 1) return nodeColors.blue
  if (level === 2) return nodeColors.green
  return nodeColors.gray
}
