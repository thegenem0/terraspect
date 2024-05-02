import { useAuth } from '@clerk/clerk-react'
import {
  useMutation,
  UseMutationResult,
  useQueryClient
} from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

type DeleteKeyParams = {
  key: string
}

const useDeleteKeyMutation = (): UseMutationResult<
  unknown,
  Error,
  DeleteKeyParams
> => {
  const { getToken } = useAuth()
  const queryClient = useQueryClient()

  const deleteKey = async (params: DeleteKeyParams): Promise<void> => {
    const api = await createAuthApi(getToken)
    return api.post('/apikey/delete', params).then((res) => res.data)
  }
  return useMutation({
    mutationKey: ['generateKey'],
    mutationFn: deleteKey,
    onError: (error: Error) => {
      console.log('error', error)
    },
    onSuccess: (data) => {
      queryClient.invalidateQueries({
        queryKey: ['keys']
      })
      return data
    }
  })
}

export { useDeleteKeyMutation }
