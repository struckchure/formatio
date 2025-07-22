import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  LoginUserMutationRequest,
  LoginUserMutationResponse,
} from "../../types/LoginUser";
import type { UseMutationOptions } from "@tanstack/react-query";

type LoginUserClient = typeof client<
  LoginUserMutationResponse,
  never,
  LoginUserMutationRequest
>;
type LoginUser = {
  data: LoginUserMutationResponse;
  error: never;
  request: LoginUserMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: LoginUserMutationResponse;
  client: {
    parameters: Partial<Parameters<LoginUserClient>[0]>;
    return: Awaited<ReturnType<LoginUserClient>>;
  };
};
/**
 * @link /auth/login/
 */
export function useLoginUser(
  options: {
    mutation?: UseMutationOptions<
      LoginUser["response"],
      LoginUser["error"],
      {
        data: LoginUser["request"];
      }
    >;
    client?: LoginUser["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        LoginUser["data"],
        LoginUser["error"],
        LoginUser["request"]
      >({
        method: "post",
        url: `/auth/login/`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
