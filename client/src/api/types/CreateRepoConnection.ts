import { TypesCreateRepoConnectionArgs } from "./types/CreateRepoConnectionArgs";
import type { DbRepoConnectionModel } from "./db/RepoConnectionModel";

/**
 * @description Created
 */
export type CreateRepoConnection201 = DbRepoConnectionModel;
/**
 * @description Create Repo Connection
 */
export type CreateRepoConnectionMutationRequest = TypesCreateRepoConnectionArgs;
/**
 * @description Created
 */
export type CreateRepoConnectionMutationResponse = DbRepoConnectionModel;
export type CreateRepoConnectionMutation = {
  Response: CreateRepoConnectionMutationResponse;
  Request: CreateRepoConnectionMutationRequest;
};
