import { useAuth } from '@clerk/clerk-react'
import type { UseQueryResult } from '@tanstack/react-query'
import { useQuery } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

export type DataNode = {
  id: string
  label: string
  variables?: VariableResponse
  children?: DataNode[]
}

export type VariableResponse = {
  simple_values: Variable[]
  complex_values: Variable[]
}

export type Variable = {
  key: string
  value: VariableValue
}

export type KVPair<TKey, TValue> = {
  key: TKey
  value: TValue
}

export type VariableValue =
  | string
  | number
  | KVPair<string, string>[]
  | string[]

export type GraphResponse = {
  tree: DataNode[]
}

export const useGraphQuery = (): UseQueryResult<GraphResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api.get('/graph').then((res) => res.data)
  }

  return useQuery({
    queryKey: ['graph'],
    queryFn: getData
  })
}
