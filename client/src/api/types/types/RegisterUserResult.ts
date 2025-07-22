import { DbCardModel } from "../db/CardModel";
import { DbGithubAccountConnectionModel } from "../db/GithubAccountConnectionModel";
import { DbInvoiceModel } from "../db/InvoiceModel";
import { DbMachineModel } from "../db/MachineModel";
import { DbSocialConnectionModel } from "../db/SocialConnectionModel";
import { DbTransactionModel } from "../db/TransactionModel";
import { LibAuthTokens } from "../lib/AuthTokens";

export type TypesRegisterUserResult = {
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
  tokens: LibAuthTokens;
  /**
   * @type string
   */
  updatedAt: string;
};
