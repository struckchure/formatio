import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  UpdateMachineMutationRequest,
  UpdateMachineMutationResponse,
  UpdateMachinePathParams,
} from "../../types/UpdateMachine";
import type { UseMutationOptions } from "@tanstack/react-query";

type UpdateMachineClient = typeof client<
  UpdateMachineMutationResponse,
  never,
  UpdateMachineMutationRequest
>;
type UpdateMachine = {
  data: UpdateMachineMutationResponse;
  error: never;
  request: UpdateMachineMutationRequest;
  pathParams: UpdateMachinePathParams;
  queryParams: never;
  headerParams: never;
  response: UpdateMachineMutationResponse;
  client: {
    parameters: Partial<Parameters<UpdateMachineClient>[0]>;
    return: Awaited<ReturnType<UpdateMachineClient>>;
  };
};
/**
 * @link /machine/:machineId
 */
export function useUpdateMachine(
  options: {
    mutation?: UseMutationOptions<
      UpdateMachine["response"],
      UpdateMachine["error"],
      {
        machineId: UpdateMachinePathParams["machineId"];
        data?: UpdateMachine["request"];
      }
    >;
    client?: UpdateMachine["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ machineId, data }) => {
      const res = await client<
        UpdateMachine["data"],
        UpdateMachine["error"],
        UpdateMachine["request"]
      >({
        method: "patch",
        url: `/machine/${machineId}`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
