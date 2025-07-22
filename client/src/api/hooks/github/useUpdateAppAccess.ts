import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type { UpdateAppAccessQueryResponse } from "../../types/UpdateAppAccess";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type UpdateAppAccessClient = typeof client<
  UpdateAppAccessQueryResponse,
  never,
  never
>;
type UpdateAppAccess = {
  data: UpdateAppAccessQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: UpdateAppAccessQueryResponse;
  client: {
    parameters: Partial<Parameters<UpdateAppAccessClient>[0]>;
    return: Awaited<ReturnType<UpdateAppAccessClient>>;
  };
};
export const updateAppAccessQueryKey = () =>
  [{ url: "/gh/update-app-access" }] as const;
export type UpdateAppAccessQueryKey = ReturnType<
  typeof updateAppAccessQueryKey
>;
export function updateAppAccessQueryOptions(
  options: UpdateAppAccess["client"]["parameters"] = {},
) {
  const queryKey = updateAppAccessQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        UpdateAppAccess["data"],
        UpdateAppAccess["error"]
      >({
        method: "get",
        url: `/gh/update-app-access`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/update-app-access
 */
export function useUpdateAppAccess<
  TData = UpdateAppAccess["response"],
  TQueryData = UpdateAppAccess["response"],
  TQueryKey extends QueryKey = UpdateAppAccessQueryKey,
>(
  options: {
    query?: Partial<
      QueryObserverOptions<
        UpdateAppAccess["response"],
        UpdateAppAccess["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: UpdateAppAccess["client"]["parameters"];
  } = {},
): UseQueryResult<TData, UpdateAppAccess["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? updateAppAccessQueryKey();
  const query = useQuery({
    ...(updateAppAccessQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, UpdateAppAccess["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const updateAppAccessSuspenseQueryKey = () =>
  [{ url: "/gh/update-app-access" }] as const;
export type UpdateAppAccessSuspenseQueryKey = ReturnType<
  typeof updateAppAccessSuspenseQueryKey
>;
export function updateAppAccessSuspenseQueryOptions(
  options: UpdateAppAccess["client"]["parameters"] = {},
) {
  const queryKey = updateAppAccessSuspenseQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        UpdateAppAccess["data"],
        UpdateAppAccess["error"]
      >({
        method: "get",
        url: `/gh/update-app-access`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/update-app-access
 */
export function useUpdateAppAccessSuspense<
  TData = UpdateAppAccess["response"],
  TQueryKey extends QueryKey = UpdateAppAccessSuspenseQueryKey,
>(
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        UpdateAppAccess["response"],
        UpdateAppAccess["error"],
        TData,
        TQueryKey
      >
    >;
    client?: UpdateAppAccess["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, UpdateAppAccess["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? updateAppAccessSuspenseQueryKey();
  const query = useSuspenseQuery({
    ...(updateAppAccessSuspenseQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, UpdateAppAccess["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
