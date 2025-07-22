import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  CreateMachinePlanMutationRequest,
  CreateMachinePlanMutationResponse,
} from "../../types/CreateMachinePlan";
import type { UseMutationOptions } from "@tanstack/react-query";

type CreateMachinePlanClient = typeof client<
  CreateMachinePlanMutationResponse,
  never,
  CreateMachinePlanMutationRequest
>;
type CreateMachinePlan = {
  data: CreateMachinePlanMutationResponse;
  error: never;
  request: CreateMachinePlanMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: CreateMachinePlanMutationResponse;
  client: {
    parameters: Partial<Parameters<CreateMachinePlanClient>[0]>;
    return: Awaited<ReturnType<CreateMachinePlanClient>>;
  };
};
/**
 * @link /plans
 */
export function useCreateMachinePlan(
  options: {
    mutation?: UseMutationOptions<
      CreateMachinePlan["response"],
      CreateMachinePlan["error"],
      {
        data: CreateMachinePlan["request"];
      }
    >;
    client?: CreateMachinePlan["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        CreateMachinePlan["data"],
        CreateMachinePlan["error"],
        CreateMachinePlan["request"]
      >({
        method: "post",
        url: `/plans`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
