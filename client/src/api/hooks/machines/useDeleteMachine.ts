import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  DeleteMachineMutationResponse,
  DeleteMachinePathParams,
} from "../../types/DeleteMachine";
import type { UseMutationOptions } from "@tanstack/react-query";

type DeleteMachineClient = typeof client<
  DeleteMachineMutationResponse,
  never,
  never
>;
type DeleteMachine = {
  data: DeleteMachineMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteMachinePathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteMachineMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteMachineClient>[0]>;
    return: Awaited<ReturnType<DeleteMachineClient>>;
  };
};
/**
 * @link /machine/:machineId
 */
export function useDeleteMachine(
  options: {
    mutation?: UseMutationOptions<
      DeleteMachine["response"],
      DeleteMachine["error"],
      {
        machineId: DeleteMachinePathParams["machineId"];
      }
    >;
    client?: DeleteMachine["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ machineId }) => {
      const res = await client<
        DeleteMachine["data"],
        DeleteMachine["error"],
        DeleteMachine["request"]
      >({
        method: "delete",
        url: `/machine/${machineId}`,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
