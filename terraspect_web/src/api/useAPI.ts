import axios from 'axios'

type ApiProps = {
  jwt: string | null
}

type GetJWT = () => Promise<string | null>

export type Api = ReturnType<typeof terraspectAPI>

const terraspectAPI = ({ jwt }: ApiProps) =>
  axios.create({
    baseURL: `http://localhost:8080/`,
    timeout: 10000,
    headers: {
      Authorization: `Bearer ${jwt}`
    }
  })

export const createAuthApi = async (jwtGetter: GetJWT) => {
  const jwt = await jwtGetter()
  if (!jwt) {
    throw new Error('User is not authenticated.')
  }
  return terraspectAPI({ jwt })
}
