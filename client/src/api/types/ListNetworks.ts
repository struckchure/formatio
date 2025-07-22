import type { DbNetworkModel } from "./db/NetworkModel";

export type ListNetworksQueryParams = {
  /**
   * @type string | undefined
   */
  machineId?: string;
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
export type ListNetworks200 = DbNetworkModel[];
/**
 * @description OK
 */
export type ListNetworksQueryResponse = DbNetworkModel[];
export type ListNetworksQuery = {
  Response: ListNetworksQueryResponse;
  QueryParams: ListNetworksQueryParams;
};
