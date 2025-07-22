import { ServicesDeployRepoArgs } from "./services/DeployRepoArgs";

/**
 * @description OK
 */
export type DeployRepo200 = unknown;
/**
 * @description Deploy Repo
 */
export type DeployRepoMutationRequest = ServicesDeployRepoArgs;
export type DeployRepoMutationResponse = unknown;
export type DeployRepoMutation = {
  Response: DeployRepoMutationResponse;
  Request: DeployRepoMutationRequest;
};
