import { useAuth } from '@clerk/clerk-react'
import { useQuery, UseQueryResult } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type ProjectsResponse = {
  projects: Project[]
}

export type Project = {
  id: string
  name: string
  description: string
  planCount: number
}

export const useAllProjectsQuery = (): UseQueryResult<ProjectsResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api.get(`/projects`).then((res) => res.data)
  }

  return useQuery({
    queryKey: ['tree'],
    queryFn: getData
  })
}
