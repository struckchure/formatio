import type { DbMachinePlanModel } from "./db/MachinePlanModel";

export type ListMachinePlansQueryParams = {
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
export type ListMachinePlans200 = DbMachinePlanModel[];
/**
 * @description OK
 */
export type ListMachinePlansQueryResponse = DbMachinePlanModel[];
export type ListMachinePlansQuery = {
  Response: ListMachinePlansQueryResponse;
  QueryParams: ListMachinePlansQueryParams;
};
