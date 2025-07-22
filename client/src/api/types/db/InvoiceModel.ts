import { DbCurrency } from "./Currency";
import { DbInvoiceStatus } from "./InvoiceStatus";
import { DbUserModel } from "./UserModel";

export type DbInvoiceModel = {
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
  from: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  productId: string;
  /**
   * @type integer
   */
  quantity: number;
  /**
   * @type string
   */
  reference: string;
  status: DbInvoiceStatus;
  /**
   * @type string
   */
  to: string;
  /**
   * @type number
   */
  totalPrice: number;
  /**
   * @type number
   */
  unitPrice: number;
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
