import { TypesRegisterUserArgs } from "./types/RegisterUserArgs";
import type { TypesRegisterUserResult } from "./types/RegisterUserResult";

/**
 * @description Created
 */
export type RegisterUser201 = TypesRegisterUserResult;
/**
 * @description Register User
 */
export type RegisterUserMutationRequest = TypesRegisterUserArgs;
/**
 * @description Created
 */
export type RegisterUserMutationResponse = TypesRegisterUserResult;
export type RegisterUserMutation = {
  Response: RegisterUserMutationResponse;
  Request: RegisterUserMutationRequest;
};
