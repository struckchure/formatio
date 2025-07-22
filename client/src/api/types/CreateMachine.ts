import { TypesCreateMachineArgs } from "./types/CreateMachineArgs";
import type { DbMachineModel } from "./db/MachineModel";

/**
 * @description Created
 */
export type CreateMachine201 = DbMachineModel[];
/**
 * @description List Machine
 */
export type CreateMachineMutationRequest = TypesCreateMachineArgs;
/**
 * @description Created
 */
export type CreateMachineMutationResponse = DbMachineModel[];
export type CreateMachineMutation = {
  Response: CreateMachineMutationResponse;
  Request: CreateMachineMutationRequest;
};
