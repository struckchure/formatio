import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  UpdateRepoConnectionMutationRequest,
  UpdateRepoConnectionMutationResponse,
  UpdateRepoConnectionPathParams,
} from "../../types/UpdateRepoConnection";
import type { UseMutationOptions } from "@tanstack/react-query";

type UpdateRepoConnectionClient = typeof client<
  UpdateRepoConnectionMutationResponse,
  never,
  UpdateRepoConnectionMutationRequest
>;
type UpdateRepoConnection = {
  data: UpdateRepoConnectionMutationResponse;
  error: never;
  request: UpdateRepoConnectionMutationRequest;
  pathParams: UpdateRepoConnectionPathParams;
  queryParams: never;
  headerParams: never;
  response: UpdateRepoConnectionMutationResponse;
  client: {
    parameters: Partial<Parameters<UpdateRepoConnectionClient>[0]>;
    return: Awaited<ReturnType<UpdateRepoConnectionClient>>;
  };
};
/**
 * @link /repo-connection/:connectionId
 */
export function useUpdateRepoConnection(
  options: {
    mutation?: UseMutationOptions<
      UpdateRepoConnection["response"],
      UpdateRepoConnection["error"],
      {
        connectionId: UpdateRepoConnectionPathParams["connectionId"];
        data?: UpdateRepoConnection["request"];
      }
    >;
    client?: UpdateRepoConnection["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ connectionId, data }) => {
      const res = await client<
        UpdateRepoConnection["data"],
        UpdateRepoConnection["error"],
        UpdateRepoConnection["request"]
      >({
        method: "patch",
        url: `/repo-connection/${connectionId}`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
