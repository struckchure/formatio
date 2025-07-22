import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeleteCardMutationResponse,
  DeleteCardPathParams,
} from "../../types/DeleteCard";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeleteCardClient = typeof client<DeleteCardMutationResponse, never, never>;
type DeleteCard = {
  data: DeleteCardMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteCardPathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteCardMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteCardClient>[0]>;
    return: Awaited<ReturnType<DeleteCardClient>>;
  };
};
/**
 * @link /billing/cards/:cardId
 */
export function useDeleteCard(
  options: {
    mutation?: UseMutationOptions<
      DeleteCard["response"],
      DeleteCard["error"],
      {
        cardId: DeleteCardPathParams["cardId"];
      }
    >;
    client?: DeleteCard["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ cardId }) => {
      const res = await client<
        DeleteCard["data"],
        DeleteCard["error"],
        DeleteCard["request"]
      >({
        method: "delete",
        url: `/billing/cards/${cardId}`,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
