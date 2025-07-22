import type { DbGithubAccountConnectionModel } from "./db/GithubAccountConnectionModel";

/**
 * @description OK
 */
export type ListAccountConnections200 = DbGithubAccountConnectionModel[];
/**
 * @description OK
 */
export type ListAccountConnectionsQueryResponse =
  DbGithubAccountConnectionModel[];
export type ListAccountConnectionsQuery = {
  Response: ListAccountConnectionsQueryResponse;
};
