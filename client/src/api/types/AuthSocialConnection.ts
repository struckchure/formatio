import { TypesAuth0UserArgs } from "./types/Auth0UserArgs";
import type { TypesLoginUserResult } from "./types/LoginUserResult";

/**
 * @description OK
 */
export type AuthSocialConnection200 = TypesLoginUserResult;
/**
 * @description Auth0 User
 */
export type AuthSocialConnectionMutationRequest = TypesAuth0UserArgs;
/**
 * @description OK
 */
export type AuthSocialConnectionMutationResponse = TypesLoginUserResult;
export type AuthSocialConnectionMutation = {
  Response: AuthSocialConnectionMutationResponse;
  Request: AuthSocialConnectionMutationRequest;
};
