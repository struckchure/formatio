import { TypesUpdateRepoConnectionArgs } from "./types/UpdateRepoConnectionArgs";
import type { DbRepoConnectionModel } from "./db/RepoConnectionModel";

export type UpdateRepoConnectionPathParams = {
  /**
   * @description Repo Connection Id
   * @type string
   */
  connectionId: string;
};
/**
 * @description Accepted
 */
export type UpdateRepoConnection202 = DbRepoConnectionModel;
/**
 * @description Update Repo Connection
 */
export type UpdateRepoConnectionMutationRequest = TypesUpdateRepoConnectionArgs;
/**
 * @description Accepted
 */
export type UpdateRepoConnectionMutationResponse = DbRepoConnectionModel;
export type UpdateRepoConnectionMutation = {
  Response: UpdateRepoConnectionMutationResponse;
  Request: UpdateRepoConnectionMutationRequest;
  PathParams: UpdateRepoConnectionPathParams;
};
