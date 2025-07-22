import { DbMachineModel } from "./MachineModel";
import { DbCurrency } from "./Currency";

export type DbMachinePlanModel = {
  /**
   * @type array
   */
  Machine: DbMachineModel[];
  /**
   * @type string
   */
  cpu: string;
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
   * @type number
   */
  hourlyRate: number;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  memory: string;
  /**
   * @type number
   */
  monthlyRate: number;
  /**
   * @type string
   */
  name: string;
  /**
   * @type string
   */
  updatedAt: string;
};
