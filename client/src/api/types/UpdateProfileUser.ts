import { TypesUpdateUserArgs } from "./types/UpdateUserArgs";
import type { DbUserModel } from "./db/UserModel";

/**
 * @description Accepted
 */
export type UpdateProfileUser202 = DbUserModel;
/**
 * @description Update User
 */
export type UpdateProfileUserMutationRequest = TypesUpdateUserArgs;
/**
 * @description Accepted
 */
export type UpdateProfileUserMutationResponse = DbUserModel;
export type UpdateProfileUserMutation = {
  Response: UpdateProfileUserMutationResponse;
  Request: UpdateProfileUserMutationRequest;
};
