import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeleteMachinePlanMutationResponse,
  DeleteMachinePlanPathParams,
} from "../../types/DeleteMachinePlan";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeleteMachinePlanClient = typeof client<
  DeleteMachinePlanMutationResponse,
  never,
  never
>;
type DeleteMachinePlan = {
  data: DeleteMachinePlanMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteMachinePlanPathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteMachinePlanMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteMachinePlanClient>[0]>;
    return: Awaited<ReturnType<DeleteMachinePlanClient>>;
  };
};
/**
 * @link /plans/:machinePlanId
 */
export function useDeleteMachinePlan(
  options: {
    mutation?: UseMutationOptions<
      DeleteMachinePlan["response"],
      DeleteMachinePlan["error"],
      {
        machinePlanId: DeleteMachinePlanPathParams["machinePlanId"];
      }
    >;
    client?: DeleteMachinePlan["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ machinePlanId }) => {
      const res = await client<
        DeleteMachinePlan["data"],
        DeleteMachinePlan["error"],
        DeleteMachinePlan["request"]
      >({
        method: "delete",
        url: `/plans/${machinePlanId}`,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
