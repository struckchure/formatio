import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  RefreshAccessTokenMutationRequest,
  RefreshAccessTokenMutationResponse,
} from "../../types/RefreshAccessToken";
import type { UseMutationOptions } from "@tanstack/react-query";

type RefreshAccessTokenClient = typeof client<
  RefreshAccessTokenMutationResponse,
  never,
  RefreshAccessTokenMutationRequest
>;
type RefreshAccessToken = {
  data: RefreshAccessTokenMutationResponse;
  error: never;
  request: RefreshAccessTokenMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: RefreshAccessTokenMutationResponse;
  client: {
    parameters: Partial<Parameters<RefreshAccessTokenClient>[0]>;
    return: Awaited<ReturnType<RefreshAccessTokenClient>>;
  };
};
/**
 * @link /auth/refresh-access-token/
 */
export function useRefreshAccessToken(
  options: {
    mutation?: UseMutationOptions<
      RefreshAccessToken["response"],
      RefreshAccessToken["error"],
      {
        data: RefreshAccessToken["request"];
      }
    >;
    client?: RefreshAccessToken["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        RefreshAccessToken["data"],
        RefreshAccessToken["error"],
        RefreshAccessToken["request"]
      >({
        method: "post",
        url: `/auth/refresh-access-token/`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
