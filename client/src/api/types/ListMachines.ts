import type { DbMachineModel } from "./db/MachineModel";

export type ListMachinesQueryParams = {
  /**
   * @type string | undefined
   */
  search?: string;
  /**
   * @type integer | undefined
   */
  skip?: number;
  /**
   * @type string | undefined
   */
  sortBy?: string;
  /**
   * @type integer | undefined
   */
  take?: number;
  /**
   * @type string | undefined
   */
  userId?: string;
};
/**
 * @description OK
 */
export type ListMachines200 = DbMachineModel[];
/**
 * @description OK
 */
export type ListMachinesQueryResponse = DbMachineModel[];
export type ListMachinesQuery = {
  Response: ListMachinesQueryResponse;
  QueryParams: ListMachinesQueryParams;
};
