import { DbDeploymentModel } from "./DeploymentModel";

export type DbDeploymentLogModel = {
  /**
   * @type string
   */
  createdAt: string;
  /**
   * @type string
   */
  deletedAt: string;
  deployment: DbDeploymentModel;
  /**
   * @type string
   */
  deploymentId: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  jobId: string;
  /**
   * @type string
   */
  message: string;
  /**
   * @type string
   */
  updatedAt: string;
};
