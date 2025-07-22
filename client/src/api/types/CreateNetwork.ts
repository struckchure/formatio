import { TypesCreateNetworkArgs } from "./types/CreateNetworkArgs";
import type { DbNetworkModel } from "./db/NetworkModel";

/**
 * @description Created
 */
export type CreateNetwork201 = DbNetworkModel;
/**
 * @description List Networks
 */
export type CreateNetworkMutationRequest = TypesCreateNetworkArgs;
/**
 * @description Created
 */
export type CreateNetworkMutationResponse = DbNetworkModel;
export type CreateNetworkMutation = {
  Response: CreateNetworkMutationResponse;
  Request: CreateNetworkMutationRequest;
};
