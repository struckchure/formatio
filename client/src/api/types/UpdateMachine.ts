import { TypesUpdateMachineArgs } from "./types/UpdateMachineArgs";

export type UpdateMachinePathParams = {
  /**
   * @description Machine Id
   * @type string
   */
  machineId: string;
};
/**
 * @description Accepted
 */
export type UpdateMachine202 = unknown;
/**
 * @description Update Machine
 */
export type UpdateMachineMutationRequest = TypesUpdateMachineArgs;
export type UpdateMachineMutationResponse = unknown;
export type UpdateMachineMutation = {
  Response: UpdateMachineMutationResponse;
  Request: UpdateMachineMutationRequest;
  PathParams: UpdateMachinePathParams;
};
