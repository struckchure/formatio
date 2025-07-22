import type { DbUserModel } from "./db/UserModel";

/**
 * @description OK
 */
export type GetProfileUser200 = DbUserModel;
/**
 * @description OK
 */
export type GetProfileUserQueryResponse = DbUserModel;
export type GetProfileUserQuery = {
  Response: GetProfileUserQueryResponse;
};
