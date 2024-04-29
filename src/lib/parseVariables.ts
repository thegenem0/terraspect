/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable @typescript-eslint/no-explicit-any */
//@ts-nocheck

import { DataNode, KVPair } from '@/hooks/useGraphQuery'

type ParsedVariables = {
  simpleValues: KVPair<string, string>[]
  objectValues: KVPair<string, KVPair<string, string>[]>[]
}

export function ParseVariables(
  activeNode?: DataNode
): ParsedVariables | undefined {
  if (!activeNode || !activeNode.variables) {
    return undefined // Return undefined if no activeNode or variables
  }

  const result: ParsedVariables = {
    simpleValues: [],
    objectValues: []
  }

  for (const variable of activeNode.variables) {
    if (
      Array.isArray(variable.value) &&
      containsOnlyObjectsWithKeyValuePairs(variable.value)
    ) {
      // This is where we handle more complex structures
      const objectPairs: KVPair<string, string>[] = variable.value.map(
        (obj) => {
          if (Object.keys(obj).length !== 2) {
            return {
              key: obj.key as string,
              value: 'Invalid object'
            }
          } else {
            return {
              key: obj.key as string,
              value: obj.value as string
            }
          }
        }
      )
      result.objectValues.push({ key: variable.key, value: objectPairs })
    } else if (Array.isArray(variable.value)) {
      // Handle simple arrays as a single string
      result.simpleValues.push({
        key: variable.key,
        value: variable.value
          .map((v) =>
            typeof v === 'object' ? JSON.stringify(v) : v.toString()
          )
          .join(', ')
      })
    } else {
      // Handle non-array values directly as simple values
      result.simpleValues.push({
        key: variable.key,
        value:
          typeof variable.value === 'object'
            ? JSON.stringify(variable.value)
            : variable.value.toString()
      })
    }
  }

  return result
}

function containsOnlyObjectsWithKeyValuePairs(array: any[]): boolean {
  return array.every(
    (item) =>
      item !== null &&
      typeof item === 'object' &&
      !Array.isArray(item) &&
      'key' in item &&
      'value' in item
  )
}
