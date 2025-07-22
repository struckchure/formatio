import { TypesUpdateMachinePlanArgs } from "./types/UpdateMachinePlanArgs";
import type { DbMachinePlanModel } from "./db/MachinePlanModel";

export type UpdateMachinePlanPathParams = {
  /**
   * @description Machine Plan Id
   * @type string
   */
  machinePlanId: string;
};
/**
 * @description Accepted
 */
export type UpdateMachinePlan202 = DbMachinePlanModel;
/**
 * @description Update Machine
 */
export type UpdateMachinePlanMutationRequest = TypesUpdateMachinePlanArgs;
/**
 * @description Accepted
 */
export type UpdateMachinePlanMutationResponse = DbMachinePlanModel;
export type UpdateMachinePlanMutation = {
  Response: UpdateMachinePlanMutationResponse;
  Request: UpdateMachinePlanMutationRequest;
  PathParams: UpdateMachinePlanPathParams;
};
