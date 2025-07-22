import { DbCurrency } from "./Currency";
import { DbTransactionStatus } from "./TransactionStatus";
import { DbTransactionType } from "./TransactionType";
import { DbUserModel } from "./UserModel";

export type DbTransactionModel = {
  /**
   * @type number
   */
  amount: number;
  /**
   * @type string
   */
  createdAt: string;
  currency: DbCurrency;
  /**
   * @type string
   */
  deletedAt: string;
  /**
   * @type string
   */
  description: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  reference: string;
  status: DbTransactionStatus;
  type: DbTransactionType;
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
