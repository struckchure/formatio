import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeleteNetworkMutationResponse,
  DeleteNetworkPathParams,
} from "../../types/DeleteNetwork";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeleteNetworkClient = typeof client<
  DeleteNetworkMutationResponse,
  never,
  never
>;
type DeleteNetwork = {
  data: DeleteNetworkMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteNetworkPathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteNetworkMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteNetworkClient>[0]>;
    return: Awaited<ReturnType<DeleteNetworkClient>>;
  };
};
/**
 * @link /network/:networkId
 */
export function useDeleteNetwork(
  options: {
    mutation?: UseMutationOptions<
      DeleteNetwork["response"],
      DeleteNetwork["error"],
      {
        networkId: DeleteNetworkPathParams["networkId"];
      }
    >;
    client?: DeleteNetwork["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ networkId }) => {
      const res = await client<
        DeleteNetwork["data"],
        DeleteNetwork["error"],
        DeleteNetwork["request"]
      >({
        method: "delete",
        url: `/network/${networkId}`,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
