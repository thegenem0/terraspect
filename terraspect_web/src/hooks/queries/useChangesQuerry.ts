import { useAuth } from '@clerk/clerk-react'
import { useQuery, UseQueryResult } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

import { ComplexVariable } from './useGraphQuery'

type ChangesResponse = {
  changes: Change[]
}

export type Change = {
  ModKey: string
  Changes: ChangeItem[]
}

export type ChangeItem = {
  Actions: string[]
  Address: string
  PreviousAddress: string
  Changes: ChangeValues
}

export type ComplexChangeVariable = {
  key: string
  value: ChangeVariableType[]
}

export type ChangeValues = {
  values: ComplexChangeVariable
}

type ChangeVariableType = {
  key: string
  value: ComplexVariable[] | string
  prev_value: ComplexVariable[] | string
}

type ChangesQueryParams = {
  projectId?: string
  planId?: string
}

export const useChangesQuery = ({
  projectId,
  planId
}: ChangesQueryParams): UseQueryResult<ChangesResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api
      .get(`/projects/${projectId}/plans/${planId}/changes`)
      .then((res) => res.data)
  }

  return useQuery({
    queryKey: ['changes', projectId, planId],
    queryFn: getData,
    enabled: !!projectId && !!planId
  })
}
