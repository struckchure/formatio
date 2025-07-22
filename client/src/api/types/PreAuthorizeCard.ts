import { TypesPreAuthorizeCardArgs } from "./types/PreAuthorizeCardArgs";
import type { TypesPreAuthorizeCardResult } from "./types/PreAuthorizeCardResult";

/**
 * @description OK
 */
export type PreAuthorizeCard200 = TypesPreAuthorizeCardResult;
/**
 * @description Pre-Authorize Card
 */
export type PreAuthorizeCardMutationRequest = TypesPreAuthorizeCardArgs;
/**
 * @description OK
 */
export type PreAuthorizeCardMutationResponse = TypesPreAuthorizeCardResult;
export type PreAuthorizeCardMutation = {
  Response: PreAuthorizeCardMutationResponse;
  Request: PreAuthorizeCardMutationRequest;
};
