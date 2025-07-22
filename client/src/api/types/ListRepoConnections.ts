import type { DbRepoConnectionModel } from "./db/RepoConnectionModel";

export type ListRepoConnectionsQueryParams = {
  /**
   * @type string | undefined
   */
  machineId?: string;
  /**
   * @type string | undefined
   */
  search?: string;
  /**
   * @type integer | undefined
   */
  skip?: number;
  /**
   * @type string | undefined
   */
  sortBy?: string;
  /**
   * @type integer | undefined
   */
  take?: number;
};
/**
 * @description OK
 */
export type ListRepoConnections200 = DbRepoConnectionModel[];
/**
 * @description OK
 */
export type ListRepoConnectionsQueryResponse = DbRepoConnectionModel[];
export type ListRepoConnectionsQuery = {
  Response: ListRepoConnectionsQueryResponse;
  QueryParams: ListRepoConnectionsQueryParams;
};
