import { useAuth } from '@clerk/clerk-react'
import {
  useMutation,
  UseMutationResult,
  useQueryClient
} from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type CreateProjectParams = {
  name: string
  description: string
}

const useCreateProjectMutation = (): UseMutationResult<
  unknown,
  Error,
  CreateProjectParams
> => {
  const { getToken } = useAuth()
  const queryClient = useQueryClient()

  const createProject = async (params: CreateProjectParams): Promise<void> => {
    const api = await createAuthApi(getToken)
    return api.post('/projects', params).then((res) => res.data)
  }
  return useMutation({
    mutationKey: ['createProject'],
    mutationFn: createProject,
    onError: (error: Error) => {
      console.log('error', error)
    },
    onSuccess: (data) => {
      queryClient.invalidateQueries({
        queryKey: ['projects']
      })
      return data
    }
  })
}

export { useCreateProjectMutation }
