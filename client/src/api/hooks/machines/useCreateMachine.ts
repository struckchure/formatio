import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  CreateMachineMutationRequest,
  CreateMachineMutationResponse,
} from "../../types/CreateMachine";
import type { UseMutationOptions } from "@tanstack/react-query";

type CreateMachineClient = typeof client<
  CreateMachineMutationResponse,
  never,
  CreateMachineMutationRequest
>;
type CreateMachine = {
  data: CreateMachineMutationResponse;
  error: never;
  request: CreateMachineMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: CreateMachineMutationResponse;
  client: {
    parameters: Partial<Parameters<CreateMachineClient>[0]>;
    return: Awaited<ReturnType<CreateMachineClient>>;
  };
};
/**
 * @link /machine
 */
export function useCreateMachine(
  options: {
    mutation?: UseMutationOptions<
      CreateMachine["response"],
      CreateMachine["error"],
      {
        data: CreateMachine["request"];
      }
    >;
    client?: CreateMachine["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        CreateMachine["data"],
        CreateMachine["error"],
        CreateMachine["request"]
      >({
        method: "post",
        url: `/machine`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
