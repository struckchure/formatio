import client from "@/api/client.ts";
import { useMutation } from "@tanstack/react-query";
import type {
  UpdateProfileUserMutationRequest,
  UpdateProfileUserMutationResponse,
} from "../../types/UpdateProfileUser";
import type { UseMutationOptions } from "@tanstack/react-query";

type UpdateProfileUserClient = typeof client<
  UpdateProfileUserMutationResponse,
  never,
  UpdateProfileUserMutationRequest
>;
type UpdateProfileUser = {
  data: UpdateProfileUserMutationResponse;
  error: never;
  request: UpdateProfileUserMutationRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: UpdateProfileUserMutationResponse;
  client: {
    parameters: Partial<Parameters<UpdateProfileUserClient>[0]>;
    return: Awaited<ReturnType<UpdateProfileUserClient>>;
  };
};
/**
 * @link /user/profile/
 */
export function useUpdateProfileUser(
  options: {
    mutation?: UseMutationOptions<
      UpdateProfileUser["response"],
      UpdateProfileUser["error"],
      {
        data: UpdateProfileUser["request"];
      }
    >;
    client?: UpdateProfileUser["client"]["parameters"];
  } = {},
) {
  const { mutation: mutationOptions, client: clientOptions = {} } =
    options ?? {};
  return useMutation({
    mutationFn: async ({ data }) => {
      const res = await client<
        UpdateProfileUser["data"],
        UpdateProfileUser["error"],
        UpdateProfileUser["request"]
      >({
        method: "patch",
        url: `/user/profile/`,
        data,
        ...clientOptions,
      });
      return res.data;
    },
    ...mutationOptions,
  });
}
