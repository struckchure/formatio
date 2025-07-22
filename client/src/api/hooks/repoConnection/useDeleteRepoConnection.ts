import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeleteRepoConnectionMutationResponse,
  DeleteRepoConnectionPathParams,
} from "../../types/DeleteRepoConnection";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeleteRepoConnectionClient = typeof client<
  DeleteRepoConnectionMutationResponse,
  never,
  never
>;
type DeleteRepoConnection = {
  data: DeleteRepoConnectionMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteRepoConnectionPathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteRepoConnectionMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteRepoConnectionClient>[0]>;
    return: Awaited<ReturnType<DeleteRepoConnectionClient>>;
  };
};
/**
 * @link /repo-connection/:connectionId
 */
export function useDeleteRepoConnection(
  options: {
    mutation?: UseMutationOptions<
      DeleteRepoConnection["response"],
      DeleteRepoConnection["error"],
      {
        connectionId: DeleteRepoConnectionPathParams["connectionId"];
      }
    >;
    client?: DeleteRepoConnection["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ connectionId }) => {
      const res = await client<
        DeleteRepoConnection["data"],
        DeleteRepoConnection["error"],
        DeleteRepoConnection["request"]
      >({
        method: "delete",
        url: `/repo-connection/${connectionId}`,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
