/* eslint-disable @typescript-eslint/no-explicit-any */
import { useTreeContext } from '@/contexts/TreeContextProvider'
import { SimpleVariable } from '@/hooks/useGraphQuery'
import { cn } from '@/lib/utils'

import DetailsDialog from '../pages/Graph/DetailsDialog/DetailsDialog'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from '../ui/accordion'

interface DetailsPanelProps {
  className?: string
}

const DetailsPanel = ({ className }: DetailsPanelProps) => {
  const { activeNode, refreshTree } = useTreeContext()

  return (
    <div
      className={cn(
        'flex w-full flex-col gap-4 text-wrap break-all px-4',
        className
      )}
    >
      <h2 className="text-2xl font-semibold text-slate-600">Details Panel</h2>
      <button onClick={() => refreshTree()}>Refresh</button>
      {activeNode ? (
        <SimpleComponent
          value={{ key: 'Resource Name', value: activeNode.label }}
        />
      ) : (
        <p>Select a node to view its properties</p>
      )}
      {activeNode && (
        <SimpleComponent
          value={{ key: 'Resource Path', value: activeNode.id }}
        />
      )}
      {activeNode &&
        activeNode.variables?.simple_values.map((variable, idx) => {
          return <SimpleComponent key={idx} value={variable} />
        })}
      {activeNode?.variables?.complex_values &&
        activeNode.variables?.complex_values.map((variable, idx) => {
          return (
            <ComplexComponent
              key={idx}
              groupName={variable.key}
              value={variable.value}
            />
          )
        })}
    </div>
  )
}

type SimpleComponentProps = {
  key?: number
  value: SimpleVariable
}

const SimpleComponent = ({ key, value }: SimpleComponentProps) => {
  return (
    <Accordion type="single" collapsible key={key}>
      <AccordionItem value={value.key}>
        <AccordionTrigger>{value.key}</AccordionTrigger>
        <AccordionContent>
          {typeof value.value === 'object' ? (
            <pre>{JSON.stringify(value.value, null, 2)}</pre>
          ) : (
            <pre>{value.value}</pre>
          )}
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}

type ComplexComponentProps = {
  key: number
  groupName: string
  value: SimpleVariable[]
}

const ComplexComponent = ({ key, groupName, value }: ComplexComponentProps) => {
  return (
    <Accordion type="single" collapsible key={key}>
      <AccordionItem value={groupName}>
        <AccordionTrigger>{groupName}</AccordionTrigger>
        <AccordionContent>
          {value.map((simpleVariable, idx) => {
            return (
              <div key={idx}>
                <h5 className="font-semibold">{simpleVariable.key}</h5>
                {typeof simpleVariable.value === 'object' ? (
                  <pre>{JSON.stringify(simpleVariable.value, null, 2)}</pre>
                ) : (
                  <pre>{simpleVariable.value}</pre>
                )}
              </div>
            )
          })}
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}

export default DetailsPanel
