import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListRepoConnectionsQueryResponse,
  ListRepoConnectionsQueryParams,
} from "../../types/ListRepoConnections";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListRepoConnectionsClient = typeof client<
  ListRepoConnectionsQueryResponse,
  never,
  never
>;
type ListRepoConnections = {
  data: ListRepoConnectionsQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListRepoConnectionsQueryParams;
  headerParams: never;
  response: ListRepoConnectionsQueryResponse;
  client: {
    parameters: Partial<Parameters<ListRepoConnectionsClient>[0]>;
    return: Awaited<ReturnType<ListRepoConnectionsClient>>;
  };
};
export const listRepoConnectionsQueryKey = (
  params?: ListRepoConnections["queryParams"],
) => [{ url: "/repo-connection" }, ...(params ? [params] : [])] as const;
export type ListRepoConnectionsQueryKey = ReturnType<
  typeof listRepoConnectionsQueryKey
>;
export function listRepoConnectionsQueryOptions(
  params?: ListRepoConnections["queryParams"],
  options: ListRepoConnections["client"]["parameters"] = {},
) {
  const queryKey = listRepoConnectionsQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListRepoConnections["data"],
        ListRepoConnections["error"]
      >({
        method: "get",
        url: `/repo-connection`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /repo-connection
 */
export function useListRepoConnections<
  TData = ListRepoConnections["response"],
  TQueryData = ListRepoConnections["response"],
  TQueryKey extends QueryKey = ListRepoConnectionsQueryKey,
>(
  params?: ListRepoConnections["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListRepoConnections["response"],
        ListRepoConnections["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListRepoConnections["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListRepoConnections["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listRepoConnectionsQueryKey(params);
  const query = useQuery({
    ...(listRepoConnectionsQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListRepoConnections["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listRepoConnectionsSuspenseQueryKey = (
  params?: ListRepoConnections["queryParams"],
) => [{ url: "/repo-connection" }, ...(params ? [params] : [])] as const;
export type ListRepoConnectionsSuspenseQueryKey = ReturnType<
  typeof listRepoConnectionsSuspenseQueryKey
>;
export function listRepoConnectionsSuspenseQueryOptions(
  params?: ListRepoConnections["queryParams"],
  options: ListRepoConnections["client"]["parameters"] = {},
) {
  const queryKey = listRepoConnectionsSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListRepoConnections["data"],
        ListRepoConnections["error"]
      >({
        method: "get",
        url: `/repo-connection`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /repo-connection
 */
export function useListRepoConnectionsSuspense<
  TData = ListRepoConnections["response"],
  TQueryKey extends QueryKey = ListRepoConnectionsSuspenseQueryKey,
>(
  params?: ListRepoConnections["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListRepoConnections["response"],
        ListRepoConnections["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListRepoConnections["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListRepoConnections["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listRepoConnectionsSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listRepoConnectionsSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListRepoConnections["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
