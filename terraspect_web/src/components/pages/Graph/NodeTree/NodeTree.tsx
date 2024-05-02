import {
  ControlsContainer,
  SigmaContainer,
  useLoadGraph,
  useRegisterEvents,
  useSetSettings,
  useSigma
} from '@react-sigma/core'
import { useLayoutCircular } from '@react-sigma/layout-circular'
import { LayoutForceAtlas2Control } from '@react-sigma/layout-forceatlas2'
import { useEffect } from 'react'

import { useTreeContext } from '@/contexts/TreeContextProvider'

import { buildGraph, GraphEdge, GraphNode } from './graphBuilder'

const SetupGraph = () => {
  const { treeData, toggleActiveNodeById, setHoveredNodeId, hoveredNodeId } =
    useTreeContext()
  const loadGraph = useLoadGraph()
  const registerEvents = useRegisterEvents()
  const setSettings = useSetSettings()
  const { assign: assignCircular } = useLayoutCircular()
  const sigma = useSigma<GraphNode, GraphEdge>()

  useEffect(() => {
    const graph = buildGraph({ data: treeData })

    loadGraph(graph)
    assignCircular()

    registerEvents({
      clickNode: (event) => toggleActiveNodeById(event.node),
      enterNode: (event) => setHoveredNodeId(event.node),
      leaveNode: () => setHoveredNodeId(undefined)
    })
  }, [
    assignCircular,
    registerEvents,
    toggleActiveNodeById,
    treeData,
    loadGraph,
    setHoveredNodeId
  ])

  useEffect(() => {
    setSettings({
      nodeReducer: (node, data) => {
        const graph = sigma.getGraph()
        const newData = {
          ...data,
          highlighted: data.highlighted || false,
          color: data.color || '#E2E2E2'
        }

        if (hoveredNodeId) {
          if (
            node === hoveredNodeId ||
            graph.neighbors(hoveredNodeId).includes(node)
          ) {
            newData.highlighted = true
          } else {
            newData.color = '#E2E2E2'
            newData.highlighted = false
          }
        }
        return newData
      },
      edgeReducer: (edge, data) => {
        const graph = sigma.getGraph()
        const newData = { ...data, hidden: false }

        if (hoveredNodeId && !graph.extremities(edge).includes(hoveredNodeId)) {
          newData.hidden = true
        }
        return newData
      }
    })
  }, [setSettings, sigma, hoveredNodeId])

  return null
}

const GraphContainer = () => {
  return (
    <div className="size-full">
      <SigmaContainer
        settings={{
          allowInvalidContainer: true
        }}
        style={{ height: '100%' }}
      >
        <SetupGraph />
        <ControlsContainer
          position="top-left"
          className="absolute left-10 top-20"
        >
          <LayoutForceAtlas2Control
            settings={{
              settings: {
                slowDown: 10,
                linLogMode: true,
                adjustSizes: true
              }
            }}
          />
        </ControlsContainer>
      </SigmaContainer>
    </div>
  )
}

export default GraphContainer