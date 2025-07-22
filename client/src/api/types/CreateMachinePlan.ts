import { TypesCreateMachinePlanArgs } from "./types/CreateMachinePlanArgs";
import type { DbMachinePlanModel } from "./db/MachinePlanModel";

/**
 * @description Created
 */
export type CreateMachinePlan201 = DbMachinePlanModel;
/**
 * @description Create Machine Plan
 */
export type CreateMachinePlanMutationRequest = TypesCreateMachinePlanArgs;
/**
 * @description Created
 */
export type CreateMachinePlanMutationResponse = DbMachinePlanModel;
export type CreateMachinePlanMutation = {
  Response: CreateMachinePlanMutationResponse;
  Request: CreateMachinePlanMutationRequest;
};
