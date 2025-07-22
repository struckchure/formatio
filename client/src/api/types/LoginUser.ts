import { TypesLoginUserArgs } from "./types/LoginUserArgs";
import type { TypesLoginUserResult } from "./types/LoginUserResult";

/**
 * @description OK
 */
export type LoginUser200 = TypesLoginUserResult;
/**
 * @description Login User
 */
export type LoginUserMutationRequest = TypesLoginUserArgs;
/**
 * @description OK
 */
export type LoginUserMutationResponse = TypesLoginUserResult;
export type LoginUserMutation = {
  Response: LoginUserMutationResponse;
  Request: LoginUserMutationRequest;
};
