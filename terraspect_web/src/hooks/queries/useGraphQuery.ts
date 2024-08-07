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
  simple_values: SimpleVariable[]
  complex_values: ComplexVariable[]
}

export type SimpleVariable = {
  key: string
  value: string | object
}

export type ComplexVariable = {
  key: string
  value: SimpleVariable[]
}

export type GraphResponse = {
  tree: {
    nodes: DataNode[]
  }
}

export type GraphQueryParams = {
  projectId?: string
  planId?: string
}

export const useGraphQuery = ({
  projectId,
  planId
}: GraphQueryParams): UseQueryResult<GraphResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api
      .get(`/projects/${projectId}/plans/${planId}/graph`)
      .then((res) => res.data)
  }

  return useQuery({
    queryKey: ['tree'],
    queryFn: getData,
    staleTime: Infinity,
    enabled: !!projectId && !!planId
  })
}
