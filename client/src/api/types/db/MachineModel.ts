import { DbDeploymentModel } from "./DeploymentModel";
import { DbNetworkModel } from "./NetworkModel";
import { DbRepoConnectionModel } from "./RepoConnectionModel";
import { DbMachineStatus } from "./MachineStatus";
import { DbUserModel } from "./UserModel";
import { DbMachinePlanModel } from "./MachinePlanModel";

export type DbMachineModel = {
  /**
   * @type array
   */
  Deployment: DbDeploymentModel[];
  /**
   * @type array
   */
  Network: DbNetworkModel[];
  /**
   * @type array
   */
  RepoConnection: DbRepoConnectionModel[];
  /**
   * @type string
   */
  containerId: string;
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
  /**
   * @type string
   */
  machineImage: string;
  /**
   * @type string
   */
  machineName: string;
  machineStatus: DbMachineStatus;
  owner: DbUserModel;
  /**
   * @type string
   */
  ownerId: string;
  plan: DbMachinePlanModel;
  /**
   * @type string
   */
  planId: string;
  /**
   * @type string
   */
  updatedAt: string;
};
