import type { DbDeploymentModel } from "./db/DeploymentModel";

export type ListDeploymentsQueryParams = {
  /**
   * @type string
   */
  machineId: string;
  /**
   * @type string | undefined
   */
  repoConnectionId?: string;
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
export type ListDeployments200 = DbDeploymentModel[];
/**
 * @description OK
 */
export type ListDeploymentsQueryResponse = DbDeploymentModel[];
export type ListDeploymentsQuery = {
  Response: ListDeploymentsQueryResponse;
  QueryParams: ListDeploymentsQueryParams;
};
