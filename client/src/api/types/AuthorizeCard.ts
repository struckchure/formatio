import { TypesAuthorizeCardArgs } from "./types/AuthorizeCardArgs";
import type { DbCardModel } from "./db/CardModel";

/**
 * @description Accepted
 */
export type AuthorizeCard202 = DbCardModel;
/**
 * @description Update Repo Connection
 */
export type AuthorizeCardMutationRequest = TypesAuthorizeCardArgs;
/**
 * @description Accepted
 */
export type AuthorizeCardMutationResponse = DbCardModel;
export type AuthorizeCardMutation = {
  Response: AuthorizeCardMutationResponse;
  Request: AuthorizeCardMutationRequest;
};
