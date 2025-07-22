import { DbUserModel } from "./UserModel";

export type DbCardModel = {
  /**
   * @type string
   */
  authToken: string;
  /**
   * @type string
   */
  cardType: string;
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
  expiryMonth: string;
  /**
   * @type string
   */
  expiryYear: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type boolean
   */
  isApproved: boolean;
  /**
   * @type boolean
   */
  isDefault: boolean;
  /**
   * @type string
   */
  lastDigits: string;
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
