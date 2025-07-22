import { TypesUpdateCardArgs } from "./types/UpdateCardArgs";
import type { DbCardModel } from "./db/CardModel";

export type UpdateCardPathParams = {
  /**
   * @description Card Id
   * @type string
   */
  cardId: string;
};
/**
 * @description Accepted
 */
export type UpdateCard202 = DbCardModel;
/**
 * @description Update Card
 */
export type UpdateCardMutationRequest = TypesUpdateCardArgs;
/**
 * @description Accepted
 */
export type UpdateCardMutationResponse = DbCardModel;
export type UpdateCardMutation = {
  Response: UpdateCardMutationResponse;
  Request: UpdateCardMutationRequest;
  PathParams: UpdateCardPathParams;
};
