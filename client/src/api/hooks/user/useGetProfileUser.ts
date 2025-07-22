import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type { GetProfileUserQueryResponse } from "../../types/GetProfileUser";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetProfileUserClient = typeof client<
  GetProfileUserQueryResponse,
  never,
  never
>;
type GetProfileUser = {
  data: GetProfileUserQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: GetProfileUserQueryResponse;
  client: {
    parameters: Partial<Parameters<GetProfileUserClient>[0]>;
    return: Awaited<ReturnType<GetProfileUserClient>>;
  };
};
export const getProfileUserQueryKey = () =>
  [{ url: "/user/profile/" }] as const;
export type GetProfileUserQueryKey = ReturnType<typeof getProfileUserQueryKey>;
export function getProfileUserQueryOptions(
  options: GetProfileUser["client"]["parameters"] = {},
) {
  const queryKey = getProfileUserQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetProfileUser["data"], GetProfileUser["error"]>(
        {
          method: "get",
          url: `/user/profile/`,
          ...options,
        },
      );
      return res.data;
    },
  });
}
/**
 * @link /user/profile/
 */
export function useGetProfileUser<
  TData = GetProfileUser["response"],
  TQueryData = GetProfileUser["response"],
  TQueryKey extends QueryKey = GetProfileUserQueryKey,
>(
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetProfileUser["response"],
        GetProfileUser["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetProfileUser["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetProfileUser["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? getProfileUserQueryKey();
  const query = useQuery({
    ...(getProfileUserQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetProfileUser["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getProfileUserSuspenseQueryKey = () =>
  [{ url: "/user/profile/" }] as const;
export type GetProfileUserSuspenseQueryKey = ReturnType<
  typeof getProfileUserSuspenseQueryKey
>;
export function getProfileUserSuspenseQueryOptions(
  options: GetProfileUser["client"]["parameters"] = {},
) {
  const queryKey = getProfileUserSuspenseQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetProfileUser["data"], GetProfileUser["error"]>(
        {
          method: "get",
          url: `/user/profile/`,
          ...options,
        },
      );
      return res.data;
    },
  });
}
/**
 * @link /user/profile/
 */
export function useGetProfileUserSuspense<
  TData = GetProfileUser["response"],
  TQueryKey extends QueryKey = GetProfileUserSuspenseQueryKey,
>(
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetProfileUser["response"],
        GetProfileUser["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetProfileUser["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetProfileUser["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? getProfileUserSuspenseQueryKey();
  const query = useSuspenseQuery({
    ...(getProfileUserSuspenseQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetProfileUser["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
