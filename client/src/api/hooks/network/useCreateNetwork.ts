import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  CreateNetworkMutationRequest,
  CreateNetworkMutationResponse,
} from "../../types/CreateNetwork";
import type { UseMutationOptions } from "@tanstack/react-query";

type CreateNetworkClient = typeof client<
  CreateNetworkMutationResponse,
  never,
  CreateNetworkMutationRequest
>;
type CreateNetwork = {
  data: CreateNetworkMutationResponse;
  error: never;
  request: CreateNetworkMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: CreateNetworkMutationResponse;
  client: {
    parameters: Partial<Parameters<CreateNetworkClient>[0]>;
    return: Awaited<ReturnType<CreateNetworkClient>>;
  };
};
/**
 * @link /network
 */
export function useCreateNetwork(
  options: {
    mutation?: UseMutationOptions<
      CreateNetwork["response"],
      CreateNetwork["error"],
      {
        data: CreateNetwork["request"];
      }
    >;
    client?: CreateNetwork["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        CreateNetwork["data"],
        CreateNetwork["error"],
        CreateNetwork["request"]
      >({
        method: "post",
        url: `/network`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
