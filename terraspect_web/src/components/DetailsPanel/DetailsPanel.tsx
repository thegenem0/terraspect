/* eslint-disable @typescript-eslint/no-explicit-any */
import { useTreeContext } from '@/contexts/TreeContextProvider'
import { cn } from '@/lib/utils'

interface DetailsPanelProps {
  className?: string
}

const DetailsPanel = ({ className }: DetailsPanelProps) => {
  const { activeNode } = useTreeContext()

  return (
    <div
      className={cn(
        'flex w-full flex-col gap-4 text-wrap break-all px-4',
        className
      )}
    >
      <h2 className="text-2xl font-semibold text-slate-600">Details Panel</h2>
      <h3>
        {activeNode
          ? `Resource Name: ${activeNode?.label}`
          : 'Select a node to view its properties'}
      </h3>
      {activeNode && <p>{`Resource Path: ${activeNode?.id}`}</p>}
      {activeNode &&
        activeNode.variables?.simple_values.map((variable, idx) => {
          return (
            <div key={idx}>
              <h4 className="text-lg font-semibold">{variable.key}</h4>
            </div>
          )
        })}
      {activeNode?.variables?.complex_values &&
        activeNode.variables?.complex_values.map((variable, idx) => {
          return (
            <div key={idx}>
              <h4 className="text-lg font-semibold">{variable.key}</h4>
            </div>
          )
        })}
    </div>
  )
}

export default DetailsPanel
