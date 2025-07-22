import { TypesRefreshAccessTokenArgs } from "./types/RefreshAccessTokenArgs";
import type { TypesLoginUserResult } from "./types/LoginUserResult";

/**
 * @description OK
 */
export type RefreshAccessToken200 = TypesLoginUserResult;
/**
 * @description Refresh Access Token
 */
export type RefreshAccessTokenMutationRequest = TypesRefreshAccessTokenArgs;
/**
 * @description OK
 */
export type RefreshAccessTokenMutationResponse = TypesLoginUserResult;
export type RefreshAccessTokenMutation = {
  Response: RefreshAccessTokenMutationResponse;
  Request: RefreshAccessTokenMutationRequest;
};
