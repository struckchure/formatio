import type { DbCardModel } from "./db/CardModel";

export type ListCardsQueryParams = {
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
export type ListCards200 = DbCardModel[];
/**
 * @description OK
 */
export type ListCardsQueryResponse = DbCardModel[];
export type ListCardsQuery = {
  Response: ListCardsQueryResponse;
  QueryParams: ListCardsQueryParams;
};
