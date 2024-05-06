import axios from 'axios'

import { env } from '@/env'

type ApiProps = {
  jwt: string | null
}

type GetJWT = () => Promise<string | null>

const terraspectAPI = ({ jwt }: ApiProps) =>
  axios.create({
    baseURL: `${env.VITE_API_BASE_URI}/api/web/v1`,
    timeout: 10000,
    headers: {
      Authorization: `Bearer ${jwt}`
    }
  })

export const createAuthApi = async (jwtGetter: GetJWT) => {
  const jwt = await jwtGetter()
  console.warn('jwt', jwt)
  if (!jwt) {
    throw new Error('User is not authenticated.')
  }
  return terraspectAPI({ jwt })
}
