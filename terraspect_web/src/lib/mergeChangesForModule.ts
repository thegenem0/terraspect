import { TreeDataNode } from '@/contexts/TreeContextProvider'
import { Change } from '@/hooks/queries/useChangesQuerry'
import { DataNode } from '@/hooks/queries/useGraphQuery'

export function mergeChangesForModule(
  dataNodes: DataNode[],
  changes: Change[]
): TreeDataNode[] {
  function traverseAndUpdateNodes(nodes: DataNode[]): TreeDataNode[] {
    return nodes.map((node) => {
      const nodeChanges = changes.filter((change) => change.ModKey === node.id)
      const changeItems = nodeChanges.flatMap((change) => change.Changes)

      const newNode: TreeDataNode = {
        ...node,
        changes: changeItems.length > 0 ? changeItems : undefined
      }

      if (node.children) {
        newNode.children = traverseAndUpdateNodes(node.children)
      }

      return newNode
    })
  }

  return traverseAndUpdateNodes(dataNodes)
}
