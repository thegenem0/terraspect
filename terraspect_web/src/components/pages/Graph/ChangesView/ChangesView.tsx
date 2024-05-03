import { ChevronRight } from 'lucide-react'

import { TreeDataNode } from '@/contexts/TreeContextProvider'
import { ChangeValues } from '@/hooks/queries/useChangesQuerry'

type ChangesViewProps = {
  activeNode?: TreeDataNode
}

const ChangesView = ({ activeNode }: ChangesViewProps) => {
  return (
    <div className="flex flex-row flex-wrap gap-6">
      {activeNode?.changes ? (
        activeNode.changes.map((change) => {
          console.log(change)
          return (
            <div key={change.Address} className="flex flex-col gap-3">
              <div className="grid grid-cols-4 gap-4">
                <div className="col-span-1">
                  <p>Address:</p>
                </div>
                <div className="col-span-3">
                  <pre className="rounded-lg bg-blue-200 p-2 text-gray-500">
                    {change.Address}
                  </pre>
                </div>
                <div className="col-span-1">
                  <p>Previous Address:</p>
                </div>
                <div className="col-span-3">
                  <pre className="rounded-lg bg-blue-200 p-2 text-gray-500">
                    {change.PreviousAddress !== ''
                      ? change.PreviousAddress
                      : 'Unchanged'}
                  </pre>
                </div>
                <div className="col-span-1">
                  <p>Actions:</p>
                </div>
                <div className="col-span-3">
                  <pre className="rounded-lg bg-blue-200 p-2 text-sm text-gray-500">
                    {change.Actions.join(', ')}
                  </pre>
                </div>
              </div>
              {change.Changes !== null && (
                <RenderChangeValues changes={change.Changes} />
              )}
            </div>
          )
        })
      ) : (
        <div className="flex flex-col gap-3">
          <p className="text-gray-500">
            This node does not have any changes associated with it.
          </p>
          <p className="text-sm text-gray-500">
            Information may be available in the child nodes.
          </p>
        </div>
      )}
    </div>
  )
}

export default ChangesView

const RenderChangeValues = ({ changes }: { changes: ChangeValues }) => {
  if (!changes || !changes.values) {
    return null
  }

  return (
    <div className="flex flex-col gap-3 rounded-lg border-2 border-black p-2">
      {changes?.values?.value?.map((variable, idx) => (
        <div key={idx} className="flex flex-col gap-1">
          <div className="flex flex-row items-center gap-2">
            <p>Variable Changed:</p>
            <pre className="rounded-lg p-2 text-black">{variable.key}</pre>
          </div>
          <div className="flex flex-row items-center gap-2 align-middle">
            {variable.value.map((value, idx) => (
              <pre
                key={idx}
                className="rounded-lg bg-red-200 p-2 text-gray-500"
              >
                {JSON.stringify(value, null, 2)}
              </pre>
            ))}
            <ChevronRight size={24} />
            {variable.prev_value.map((value, idx) => (
              <div key={idx} className="flex flex-row gap-1">
                <pre className="rounded-lg bg-green-200 p-2 text-gray-500">
                  {JSON.stringify(value, null, 2)}
                </pre>
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  )
}
