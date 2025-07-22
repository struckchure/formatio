import type { DbMachineModel } from "./db/MachineModel";

export type GetMachinePathParams = {
  /**
   * @description Machine Id
   * @type string
   */
  machineId: string;
};
/**
 * @description OK
 */
export type GetMachine200 = DbMachineModel;
/**
 * @description OK
 */
export type GetMachineQueryResponse = DbMachineModel;
export type GetMachineQuery = {
  Response: GetMachineQueryResponse;
  PathParams: GetMachinePathParams;
};
