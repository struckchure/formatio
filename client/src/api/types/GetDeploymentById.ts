import type { DbDeploymentModel } from "./db/DeploymentModel";

export type GetDeploymentByIdPathParams = {
  /**
   * @description Deployment Id
   * @type string
   */
  deploymentId: string;
};
/**
 * @description OK
 */
export type GetDeploymentById200 = DbDeploymentModel;
/**
 * @description OK
 */
export type GetDeploymentByIdQueryResponse = DbDeploymentModel;
export type GetDeploymentByIdQuery = {
  Response: GetDeploymentByIdQueryResponse;
  PathParams: GetDeploymentByIdPathParams;
};
