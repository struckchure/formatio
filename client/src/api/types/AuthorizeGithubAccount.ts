import type { HandlersRedirectResult } from "./handlers/RedirectResult";

export type AuthorizeGithubAccountQueryParams = {
  /**
   * @type string
   */
  redirectUrl: string;
};
/**
 * @description OK
 */
export type AuthorizeGithubAccount200 = HandlersRedirectResult;
/**
 * @description OK
 */
export type AuthorizeGithubAccountQueryResponse = HandlersRedirectResult;
export type AuthorizeGithubAccountQuery = {
  Response: AuthorizeGithubAccountQueryResponse;
  QueryParams: AuthorizeGithubAccountQueryParams;
};
