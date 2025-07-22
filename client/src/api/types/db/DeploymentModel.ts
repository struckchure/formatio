import { DbDeploymentLogModel } from "./DeploymentLogModel";
import { DbMachineModel } from "./MachineModel";
import { DbRepoConnectionModel } from "./RepoConnectionModel";
import { DbDeploymentStatus } from "./DeploymentStatus";

export type DbDeploymentModel = {
  /**
   * @type array
   */
  DeploymentLog: DbDeploymentLogModel[];
  /**
   * @type string
   */
  actor: string;
  /**
   * @type string
   */
  commitHash: string;
  /**
   * @type string
   */
  commitMessage: string;
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
  repoConnection: DbRepoConnectionModel;
  /**
   * @type string
   */
  repoConnectionId: string;
  status: DbDeploymentStatus;
  /**
   * @type string
   */
  updatedAt: string;
};
