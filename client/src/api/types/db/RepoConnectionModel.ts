import { DbDeploymentModel } from "./DeploymentModel";
import { DbMachineModel } from "./MachineModel";

export type DbRepoConnectionModel = {
  /**
   * @type array
   */
  Deployment: DbDeploymentModel[];
  /**
   * @type string
   */
  createdAt: string;
  /**
   * @type string
   */
  deletedAt: string;
  /**
   * @type string
   */
  id: string;
  machine: DbMachineModel;
  /**
   * @type string
   */
  machineId: string;
  /**
   * @type string
   */
  repoId: string;
  /**
   * @type string
   */
  repoName: string;
  /**
   * @type string
   */
  updatedAt: string;
};
