import { useAuth } from '@clerk/clerk-react'
import { useQuery, UseQueryResult } from '@tanstack/react-query'

import { createAuthApi } from '@/api/useAPI'

export type Key = {
  id: string
  name: string
  description: string
  key: string
  created_at: string
}

type KeysResponse = {
  keys: Key[]
}

export const useGetKeysQuery = (): UseQueryResult<KeysResponse> => {
  const { getToken } = useAuth()

  const getData = async () => {
    const api = await createAuthApi(getToken)
    return api.get('/apikey').then((res) => res.data)
  }

  return useQuery({
    queryKey: ['keys'],
    queryFn: getData
  })
}
