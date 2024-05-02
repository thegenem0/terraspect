import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger
} from '@/components/ui/accordion'
import { DataNode, SimpleVariable } from '@/hooks/queries/useGraphQuery'

type DetailsViewProps = {
  activeNode?: DataNode
}

const DetailsView = ({ activeNode }: DetailsViewProps) => {
  const { simple_values, complex_values } = activeNode?.variables ?? {
    simple_values: [],
    complex_values: []
  }

  return (
    <div className="flex flex-row flex-wrap gap-6">
      {!simple_values.length && !complex_values.length ? (
        <EmptyComponent />
      ) : (
        <>
          <div className="flex-1">
            {simple_values.map((variable, idx) => (
              <SimpleComponent key={idx} value={variable} />
            ))}
          </div>
          <div className="flex-1">
            {complex_values?.length > 0 &&
              complex_values.map((variable, idx) => (
                <ComplexComponent
                  key={idx}
                  groupName={variable.key}
                  value={variable.value}
                />
              ))}
          </div>
        </>
      )}
    </div>
  )
}

export default DetailsView

const EmptyComponent = () => {
  return (
    <div className="flex flex-col gap-3">
      <p className="text-gray-500">
        This node does not have any variables associated with it.
      </p>
      <p className="text-sm text-gray-500">
        Information may be available in the child nodes.
      </p>
    </div>
  )
}

type SimpleComponentProps = {
  key?: number
  value: SimpleVariable
}

const SimpleComponent = ({ key, value }: SimpleComponentProps) => {
  return (
    <Accordion type="single" collapsible key={key} className="max-w-sm">
      <AccordionItem value={value.key}>
        <AccordionTrigger>{value.key}</AccordionTrigger>
        <AccordionContent>
          {typeof value.value === 'object' ? (
            <pre>{JSON.stringify(value.value, null, 2)}</pre>
          ) : (
            <pre className="text-wrap">
              {typeof value.value === 'boolean'
                ? value.value
                  ? 'true'
                  : 'false'
                : value.value}
            </pre>
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
    <Accordion type="single" collapsible key={key} className="max-w-sm">
      <AccordionItem value={groupName}>
        <AccordionTrigger>{groupName}</AccordionTrigger>
        <AccordionContent>
          {value.map((simpleVariable, idx) => {
            return (
              <div key={idx}>
                <h5 className="font-semibold">{simpleVariable.key}</h5>
                {typeof simpleVariable.value === 'object' ? (
                  <pre className="text-wrap">
                    {JSON.stringify(simpleVariable.value, null, 2)}
                  </pre>
                ) : (
                  <pre className="text-wrap">
                    {typeof simpleVariable.value === 'boolean'
                      ? simpleVariable.value
                        ? 'true'
                        : 'false'
                      : simpleVariable.value}
                  </pre>
                )}
              </div>
            )
          })}
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}
