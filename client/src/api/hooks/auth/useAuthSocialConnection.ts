import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  AuthSocialConnectionMutationRequest,
  AuthSocialConnectionMutationResponse,
} from "../../types/AuthSocialConnection";
import type { UseMutationOptions } from "@tanstack/react-query";

type AuthSocialConnectionClient = typeof client<
  AuthSocialConnectionMutationResponse,
  never,
  AuthSocialConnectionMutationRequest
>;
type AuthSocialConnection = {
  data: AuthSocialConnectionMutationResponse;
  error: never;
  request: AuthSocialConnectionMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: AuthSocialConnectionMutationResponse;
  client: {
    parameters: Partial<Parameters<AuthSocialConnectionClient>[0]>;
    return: Awaited<ReturnType<AuthSocialConnectionClient>>;
  };
};
/**
 * @link /auth/social-connection/
 */
export function useAuthSocialConnection(
  options: {
    mutation?: UseMutationOptions<
      AuthSocialConnection["response"],
      AuthSocialConnection["error"],
      {
        data: AuthSocialConnection["request"];
      }
    >;
    client?: AuthSocialConnection["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        AuthSocialConnection["data"],
        AuthSocialConnection["error"],
        AuthSocialConnection["request"]
      >({
        method: "post",
        url: `/auth/social-connection/`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
