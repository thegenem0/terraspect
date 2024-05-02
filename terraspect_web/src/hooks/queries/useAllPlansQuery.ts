import { useAuth } from '@clerk/clerk-react'
import { useQuery, UseQueryResult } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type PlansResponse = {
  plans: Plan[]
}

export type Plan = {
  id: string
  projectId: string
  createdAt: string
}

type PlanQueryParams = {
  projectId?: string
}

export const useAllPlansQuery = ({
  projectId
}: PlanQueryParams): UseQueryResult<PlansResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api.get(`/projects/${projectId}/plans`).then((res) => res.data)
  }

  return useQuery({
    queryKey: ['plans'],
    queryFn: getData,
    enabled: !!projectId
  })
}
