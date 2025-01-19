import { userManager } from './oidc-config'
import { isTokenExpired, refreshTokens } from './token-utils'

export async function login() {
  localStorage.removeItem('loggedInUser')
  await userManager.signinRedirect()
}

export function logout() {
  localStorage.removeItem('loggedInUser')
  window.location.href = `${import.meta.env.VITE_COGNITO_DOMAIN}/logout?client_id=${import.meta.env.VITE_COGNITO_CLIENT_ID}&logout_uri=${encodeURIComponent('http://localhost:5173/')}`
}

export async function getAccessToken(): Promise<string | null> {
  let accessToken = localStorage.getItem('access_token')
  let idToken = localStorage.getItem('id_token')
  let refreshToken = localStorage.getItem('refresh_token')

  if (accessToken && isTokenExpired(accessToken)) {
    if (refreshToken) {
      const updatedUser = await refreshTokens(refreshToken)
      if (updatedUser) {
        accessToken = updatedUser.access_token
        idToken = updatedUser.id_token
        if (accessToken && idToken) {
          localStorage.setItem('access_token', accessToken)
          localStorage.setItem('id_token', idToken)
        } else {
          throw new Error('Updated access token is null')
        }
      } else {
        throw new Error('Failed to refresh access token')
      }
    }
  }
  return accessToken
}
