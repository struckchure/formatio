import type { DbDeploymentLogModel } from "./db/DeploymentLogModel";

export type ListDeploymentLogsByIdPathParams = {
  /**
   * @description Deployment Id
   * @type string
   */
  deploymentId: string;
};
/**
 * @description OK
 */
export type ListDeploymentLogsById200 = DbDeploymentLogModel[];
/**
 * @description OK
 */
export type ListDeploymentLogsByIdQueryResponse = DbDeploymentLogModel[];
export type ListDeploymentLogsByIdQuery = {
  Response: ListDeploymentLogsByIdQueryResponse;
  PathParams: ListDeploymentLogsByIdPathParams;
};
