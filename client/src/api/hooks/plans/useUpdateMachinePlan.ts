import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  UpdateMachinePlanMutationRequest,
  UpdateMachinePlanMutationResponse,
  UpdateMachinePlanPathParams,
} from "../../types/UpdateMachinePlan";
import type { UseMutationOptions } from "@tanstack/react-query";

type UpdateMachinePlanClient = typeof client<
  UpdateMachinePlanMutationResponse,
  never,
  UpdateMachinePlanMutationRequest
>;
type UpdateMachinePlan = {
  data: UpdateMachinePlanMutationResponse;
  error: never;
  request: UpdateMachinePlanMutationRequest;
  pathParams: UpdateMachinePlanPathParams;
  queryParams: never;
  headerParams: never;
  response: UpdateMachinePlanMutationResponse;
  client: {
    parameters: Partial<Parameters<UpdateMachinePlanClient>[0]>;
    return: Awaited<ReturnType<UpdateMachinePlanClient>>;
  };
};
/**
 * @link /plans/:machinePlanId
 */
export function useUpdateMachinePlan(
  options: {
    mutation?: UseMutationOptions<
      UpdateMachinePlan["response"],
      UpdateMachinePlan["error"],
      {
        machinePlanId: UpdateMachinePlanPathParams["machinePlanId"];
        data?: UpdateMachinePlan["request"];
      }
    >;
    client?: UpdateMachinePlan["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ machinePlanId, data }) => {
      const res = await client<
        UpdateMachinePlan["data"],
        UpdateMachinePlan["error"],
        UpdateMachinePlan["request"]
      >({
        method: "patch",
        url: `/plans/${machinePlanId}`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
