import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  RegisterUserMutationRequest,
  RegisterUserMutationResponse,
} from "../../types/RegisterUser";
import type { UseMutationOptions } from "@tanstack/react-query";

type RegisterUserClient = typeof client<
  RegisterUserMutationResponse,
  never,
  RegisterUserMutationRequest
>;
type RegisterUser = {
  data: RegisterUserMutationResponse;
  error: never;
  request: RegisterUserMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: RegisterUserMutationResponse;
  client: {
    parameters: Partial<Parameters<RegisterUserClient>[0]>;
    return: Awaited<ReturnType<RegisterUserClient>>;
  };
};
/**
 * @link /auth/register/
 */
export function useRegisterUser(
  options: {
    mutation?: UseMutationOptions<
      RegisterUser["response"],
      RegisterUser["error"],
      {
        data: RegisterUser["request"];
      }
    >;
    client?: RegisterUser["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        RegisterUser["data"],
        RegisterUser["error"],
        RegisterUser["request"]
      >({
        method: "post",
        url: `/auth/register/`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
