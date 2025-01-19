import axios from 'axios'

export function isTokenExpired(token: string): boolean {
  const payload = JSON.parse(atob(token.split('.')[1]))
  const exp = payload.exp * 1000
  return Date.now() > exp
}

export async function refreshTokens(refreshToken: string): Promise<any> {
  try {
    const response = await axios.post(
      `${import.meta.env.VITE_COGNITO_DOMAIN}/oauth2/token`,
      new URLSearchParams({
        grant_type: 'refresh_token',
        refresh_token: refreshToken,
        client_id: import.meta.env.VITE_COGNITO_CLIENT_ID,
      }),
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      },
    )

    if (response.data.access_token && response.data.id_token) {
      const updatedUser = JSON.parse(localStorage.getItem('loggedInUser') || '{}')
      updatedUser.id_token = response.data.id_token
      updatedUser.access_token = response.data.access_token
      return updatedUser
    } else {
      throw new Error('Missing access token or id token in response')
    }
  } catch (error) {
    console.error('Error refreshing tokens:', error)
    throw new Error('Failed to refresh tokens. Please log in again.')
  }
}
