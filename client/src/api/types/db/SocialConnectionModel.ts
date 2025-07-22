import { DbUserModel } from "./UserModel";

export type DbSocialConnectionModel = {
  /**
   * @type string
   */
  connectionId: string;
  /**
   * @type string
   */
  connectionType: string;
  /**
   * @type string
   */
  createdAt: string;
  /**
   * @type string
   */
  deletedAt: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  updatedAt: string;
  user: DbUserModel;
  /**
   * @type string
   */
  userId: string;
};
