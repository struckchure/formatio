import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  AuthorizeCardMutationRequest,
  AuthorizeCardMutationResponse,
} from "../../types/AuthorizeCard";
import type { UseMutationOptions } from "@tanstack/react-query";

type AuthorizeCardClient = typeof client<
  AuthorizeCardMutationResponse,
  never,
  AuthorizeCardMutationRequest
>;
type AuthorizeCard = {
  data: AuthorizeCardMutationResponse;
  error: never;
  request: AuthorizeCardMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: AuthorizeCardMutationResponse;
  client: {
    parameters: Partial<Parameters<AuthorizeCardClient>[0]>;
    return: Awaited<ReturnType<AuthorizeCardClient>>;
  };
};
/**
 * @link /billing/cards/authorize
 */
export function useAuthorizeCard(
  options: {
    mutation?: UseMutationOptions<
      AuthorizeCard["response"],
      AuthorizeCard["error"],
      {
        data: AuthorizeCard["request"];
      }
    >;
    client?: AuthorizeCard["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        AuthorizeCard["data"],
        AuthorizeCard["error"],
        AuthorizeCard["request"]
      >({
        method: "post",
        url: `/billing/cards/authorize`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
