import { useAuth } from '@clerk/clerk-react'
import {
  useMutation,
  UseMutationResult,
  useQueryClient
} from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type DeleteProjectParams = {
  projectId: string
}

const useDeleteProjectMutation = (): UseMutationResult<
  unknown,
  Error,
  DeleteProjectParams
> => {
  const { getToken } = useAuth()
  const queryClient = useQueryClient()

  const deleteProject = async ({
    projectId
  }: DeleteProjectParams): Promise<void> => {
    const api = await createAuthApi(getToken)
    return api.delete(`/projects/${projectId}`).then((res) => res.data)
  }
  return useMutation({
    mutationKey: ['deleteProject'],
    mutationFn: deleteProject,
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

export { useDeleteProjectMutation }
