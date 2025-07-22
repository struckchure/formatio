import type { DbRepoConnectionModel } from "./db/RepoConnectionModel";

export type GetRepoConnectionPathParams = {
  /**
   * @description Repo Connection Id
   * @type string
   */
  connectionId: string;
};
/**
 * @description No Content
 */
export type GetRepoConnection204 = DbRepoConnectionModel;
/**
 * @description No Content
 */
export type GetRepoConnectionQueryResponse = DbRepoConnectionModel;
export type GetRepoConnectionQuery = {
  Response: GetRepoConnectionQueryResponse;
  PathParams: GetRepoConnectionPathParams;
};
