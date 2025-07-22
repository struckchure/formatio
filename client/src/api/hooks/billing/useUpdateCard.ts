import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  UpdateCardMutationRequest,
  UpdateCardMutationResponse,
  UpdateCardPathParams,
} from "../../types/UpdateCard";
import type { UseMutationOptions } from "@tanstack/react-query";

type UpdateCardClient = typeof client<
  UpdateCardMutationResponse,
  never,
  UpdateCardMutationRequest
>;
type UpdateCard = {
  data: UpdateCardMutationResponse;
  error: never;
  request: UpdateCardMutationRequest;
  pathParams: UpdateCardPathParams;
  queryParams: never;
  headerParams: never;
  response: UpdateCardMutationResponse;
  client: {
    parameters: Partial<Parameters<UpdateCardClient>[0]>;
    return: Awaited<ReturnType<UpdateCardClient>>;
  };
};
/**
 * @link /billing/cards/:cardId
 */
export function useUpdateCard(
  options: {
    mutation?: UseMutationOptions<
      UpdateCard["response"],
      UpdateCard["error"],
      {
        cardId: UpdateCardPathParams["cardId"];
        data?: UpdateCard["request"];
      }
    >;
    client?: UpdateCard["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ cardId, data }) => {
      const res = await client<
        UpdateCard["data"],
        UpdateCard["error"],
        UpdateCard["request"]
      >({
        method: "patch",
        url: `/billing/cards/${cardId}`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
