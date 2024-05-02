import { useAuth } from '@clerk/clerk-react'
import { useQuery, UseQueryResult } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type ProjectResponse = {
  project: {
    id: string
    name: string
    description: string
  }
}

type ProjectQueryParams = {
  projectId?: string
}

export const useProjectQuery = ({
  projectId
}: ProjectQueryParams): UseQueryResult<ProjectResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api.get(`/projects/${projectId}`).then((res) => res.data)
  }

  return useQuery({
    queryKey: ['project', projectId],
    queryFn: getData,
    enabled: !!projectId
  })
}
