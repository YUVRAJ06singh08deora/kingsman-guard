export const cognitoAuthConfig = {
  authority: import.meta.env.VITE_COGNITO_AUTHORITY,
  client_id: import.meta.env.VITE_COGNITO_CLIENT_ID,
  redirect_uri: 'http://localhost:5173/authenticate',
  response_type: 'code',
  scope: 'email openid aws.cognito.signin.user.admin',
  token_endpoint: `${import.meta.env.VITE_COGNITO_DOMAIN}/oauth2/token`,
}
