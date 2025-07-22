import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeployRepoMutationRequest,
  DeployRepoMutationResponse,
} from "../../types/DeployRepo";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeployRepoClient = typeof client<
  DeployRepoMutationResponse,
  never,
  DeployRepoMutationRequest
>;
type DeployRepo = {
  data: DeployRepoMutationResponse;
  error: never;
  request: DeployRepoMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: DeployRepoMutationResponse;
  client: {
    parameters: Partial<Parameters<DeployRepoClient>[0]>;
    return: Awaited<ReturnType<DeployRepoClient>>;
  };
};
/**
 * @link /deployments/deploy
 */
export function useDeployRepo(
  options: {
    mutation?: UseMutationOptions<
      DeployRepo["response"],
      DeployRepo["error"],
      {
        data: DeployRepo["request"];
      }
    >;
    client?: DeployRepo["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        DeployRepo["data"],
        DeployRepo["error"],
        DeployRepo["request"]
      >({
        method: "post",
        url: `/deployments/deploy`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
