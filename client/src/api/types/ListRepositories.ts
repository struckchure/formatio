import type { TypesRepository } from "./types/Repository";

export type ListRepositoriesQueryParams = {
  /**
   * @type integer | undefined
   */
  pageNumber?: number;
  /**
   * @type integer | undefined
   */
  pageSize?: number;
};
/**
 * @description OK
 */
export type ListRepositories200 = TypesRepository[];
/**
 * @description OK
 */
export type ListRepositoriesQueryResponse = TypesRepository[];
export type ListRepositoriesQuery = {
  Response: ListRepositoriesQueryResponse;
  QueryParams: ListRepositoriesQueryParams;
};
