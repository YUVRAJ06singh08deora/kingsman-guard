import { UserManager } from 'oidc-client-ts'
import { cognitoAuthConfig } from './config'

export const userManager = new UserManager({
  ...cognitoAuthConfig,
})
