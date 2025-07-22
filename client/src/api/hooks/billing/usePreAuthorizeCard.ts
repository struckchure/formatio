import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  PreAuthorizeCardMutationRequest,
  PreAuthorizeCardMutationResponse,
} from "../../types/PreAuthorizeCard";
import type { UseMutationOptions } from "@tanstack/react-query";

type PreAuthorizeCardClient = typeof client<
  PreAuthorizeCardMutationResponse,
  never,
  PreAuthorizeCardMutationRequest
>;
type PreAuthorizeCard = {
  data: PreAuthorizeCardMutationResponse;
  error: never;
  request: PreAuthorizeCardMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: PreAuthorizeCardMutationResponse;
  client: {
    parameters: Partial<Parameters<PreAuthorizeCardClient>[0]>;
    return: Awaited<ReturnType<PreAuthorizeCardClient>>;
  };
};
/**
 * @link /billing/cards/pre-authorize
 */
export function usePreAuthorizeCard(
  options: {
    mutation?: UseMutationOptions<
      PreAuthorizeCard["response"],
      PreAuthorizeCard["error"],
      {
        data: PreAuthorizeCard["request"];
      }
    >;
    client?: PreAuthorizeCard["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        PreAuthorizeCard["data"],
        PreAuthorizeCard["error"],
        PreAuthorizeCard["request"]
      >({
        method: "post",
        url: `/billing/cards/pre-authorize`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
