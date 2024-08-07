import {
  ControlsContainer,
  SearchControl,
  SigmaContainer,
  useLoadGraph,
  useRegisterEvents,
  useSetSettings,
  useSigma,
  ZoomControl
} from '@react-sigma/core'
import { useLayoutCircular } from '@react-sigma/layout-circular'
import { LayoutForceAtlas2Control } from '@react-sigma/layout-forceatlas2'
import { useEffect } from 'react'

import { Button } from '@/components/ui/button'
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

const SigmaActions = () => {
  const { assign: assignCircular } = useLayoutCircular()

  return (
    <ControlsContainer position="top-left" className="absolute left-20 top-28">
      <div className="flex flex-row items-center gap-2">
        <Button variant="destructive" onClick={() => assignCircular()}>
          Reset
        </Button>
        <ZoomControl />
        <LayoutForceAtlas2Control
          settings={{
            settings: {
              slowDown: 10,
              linLogMode: true,
              adjustSizes: true,
              outboundAttractionDistribution: true
            }
          }}
        />
      </div>
    </ControlsContainer>
  )
}

const LegendContaier = () => {
  return (
    <div className="absolute bottom-20 left-20">
      <div className="flex flex-col gap-2">
        <div className="flex flex-row items-center gap-2">
          <div className="size-4 rounded-full bg-ts-purple"></div>
          <span>Root Node</span>
        </div>
        <div className="flex flex-row items-center gap-2">
          <div className="size-4 rounded-full bg-ts-blue"></div>
          <span>Module Node</span>
        </div>
        <div className="flex flex-row items-center gap-2">
          <div className="size-4 rounded-full bg-ts-green"></div>
          <span>Module Child Node</span>
        </div>
        <div className="flex flex-row items-center gap-2">
          <div className="size-4 rounded-full bg-ts-dark-gray"></div>
          <span>Sub-Module Node</span>
        </div>
        <div className="flex flex-row items-center gap-2">
          <div className="size-4 rounded-full bg-ts-light-red"></div>
          <span>Node with changes</span>
        </div>
      </div>
    </div>
  )
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
        <SigmaActions />
        <LegendContaier />
        <ControlsContainer
          position="top-right"
          className="absolute right-20 top-28"
        >
          <SearchControl className="overflow-hidden rounded-lg border-2 border-black bg-white px-2 py-1 outline-none" />
        </ControlsContainer>
      </SigmaContainer>
    </div>
  )
}

export default GraphContainer
