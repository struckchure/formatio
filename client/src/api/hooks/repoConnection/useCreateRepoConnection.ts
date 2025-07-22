import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  CreateRepoConnectionMutationRequest,
  CreateRepoConnectionMutationResponse,
} from "../../types/CreateRepoConnection";
import type { UseMutationOptions } from "@tanstack/react-query";

type CreateRepoConnectionClient = typeof client<
  CreateRepoConnectionMutationResponse,
  never,
  CreateRepoConnectionMutationRequest
>;
type CreateRepoConnection = {
  data: CreateRepoConnectionMutationResponse;
  error: never;
  request: CreateRepoConnectionMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: CreateRepoConnectionMutationResponse;
  client: {
    parameters: Partial<Parameters<CreateRepoConnectionClient>[0]>;
    return: Awaited<ReturnType<CreateRepoConnectionClient>>;
  };
};
/**
 * @link /repo-connection
 */
export function useCreateRepoConnection(
  options: {
    mutation?: UseMutationOptions<
      CreateRepoConnection["response"],
      CreateRepoConnection["error"],
      {
        data: CreateRepoConnection["request"];
      }
    >;
    client?: CreateRepoConnection["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        CreateRepoConnection["data"],
        CreateRepoConnection["error"],
        CreateRepoConnection["request"]
      >({
        method: "post",
        url: `/repo-connection`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
