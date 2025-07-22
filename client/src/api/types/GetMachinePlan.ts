import type { DbMachinePlanModel } from "./db/MachinePlanModel";

export type GetMachinePlanPathParams = {
  /**
   * @description Machine Plan Id
   * @type string
   */
  machinePlanId: string;
};
/**
 * @description OK
 */
export type GetMachinePlan200 = DbMachinePlanModel;
/**
 * @description OK
 */
export type GetMachinePlanQueryResponse = DbMachinePlanModel;
export type GetMachinePlanQuery = {
  Response: GetMachinePlanQueryResponse;
  PathParams: GetMachinePlanPathParams;
};
