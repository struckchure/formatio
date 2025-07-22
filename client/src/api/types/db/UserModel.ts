import { DbCardModel } from "./CardModel";
import { DbGithubAccountConnectionModel } from "./GithubAccountConnectionModel";
import { DbInvoiceModel } from "./InvoiceModel";
import { DbMachineModel } from "./MachineModel";
import { DbSocialConnectionModel } from "./SocialConnectionModel";
import { DbTransactionModel } from "./TransactionModel";

export type DbUserModel = {
  /**
   * @type array
   */
  Card: DbCardModel[];
  /**
   * @type array
   */
  GithubAccountConnection: DbGithubAccountConnectionModel[];
  /**
   * @type array
   */
  Invoice: DbInvoiceModel[];
  /**
   * @type array
   */
  Machine: DbMachineModel[];
  /**
   * @type array
   */
  SocialConnection: DbSocialConnectionModel[];
  /**
   * @type array
   */
  Transaction: DbTransactionModel[];
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
  email: string;
  /**
   * @type string
   */
  firstName: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  lastName: string;
  /**
   * @type string
   */
  password: string;
  /**
   * @type array
   */
  roles: string[];
  /**
   * @type string
   */
  updatedAt: string;
};
