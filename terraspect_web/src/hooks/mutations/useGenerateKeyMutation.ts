import { useAuth } from '@clerk/clerk-react'
import {
  useMutation,
  UseMutationResult,
  useQueryClient
} from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

export type GenerateKeyResponse = {
  key: string
}

type GenerateKeyParams = {
  name: string
  description: string
}

const useGenerateKeyMutation = (): UseMutationResult<
  GenerateKeyResponse,
  Error,
  GenerateKeyParams
> => {
  const { getToken } = useAuth()
  const queryClient = useQueryClient()

  const generateKey = async (
    params: GenerateKeyParams
  ): Promise<GenerateKeyResponse> => {
    const api = await createAuthApi(getToken)
    return api
      .post<GenerateKeyResponse>('/apikey', params)
      .then((res) => res.data)
  }
  return useMutation({
    mutationKey: ['generateKey'],
    mutationFn: generateKey,
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

export { useGenerateKeyMutation }
