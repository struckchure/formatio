import { DbMachineModel } from "./MachineModel";

export type DbNetworkModel = {
  /**
   * @type string
   */
  createdAt: string;
  /**
   * @type string
   */
  deletedAt: string;
  /**
   * @type integer
   */
  destinationPort: number;
  /**
   * @type string
   */
  hostName: string;
  /**
   * @type string
   */
  id: string;
  /**
   * @type string
   */
  ingressId: string;
  /**
   * @type integer
   */
  listeningPort: number;
  machine: DbMachineModel;
  /**
   * @type string
   */
  machineId: string;
  /**
   * @type string
   */
  protocol: string;
  /**
   * @type string
   */
  serviceId: string;
  /**
   * @type string
   */
  updatedAt: string;
};
