import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type { ListAccountConnectionsQueryResponse } from "../../types/ListAccountConnections";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListAccountConnectionsClient = typeof client<
  ListAccountConnectionsQueryResponse,
  never,
  never
>;
type ListAccountConnections = {
  data: ListAccountConnectionsQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: ListAccountConnectionsQueryResponse;
  client: {
    parameters: Partial<Parameters<ListAccountConnectionsClient>[0]>;
    return: Awaited<ReturnType<ListAccountConnectionsClient>>;
  };
};
export const listAccountConnectionsQueryKey = () =>
  [{ url: "/gh/account-connections" }] as const;
export type ListAccountConnectionsQueryKey = ReturnType<
  typeof listAccountConnectionsQueryKey
>;
export function listAccountConnectionsQueryOptions(
  options: ListAccountConnections["client"]["parameters"] = {},
) {
  const queryKey = listAccountConnectionsQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListAccountConnections["data"],
        ListAccountConnections["error"]
      >({
        method: "get",
        url: `/gh/account-connections`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/account-connections
 */
export function useListAccountConnections<
  TData = ListAccountConnections["response"],
  TQueryData = ListAccountConnections["response"],
  TQueryKey extends QueryKey = ListAccountConnectionsQueryKey,
>(
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListAccountConnections["response"],
        ListAccountConnections["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListAccountConnections["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListAccountConnections["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listAccountConnectionsQueryKey();
  const query = useQuery({
    ...(listAccountConnectionsQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListAccountConnections["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listAccountConnectionsSuspenseQueryKey = () =>
  [{ url: "/gh/account-connections" }] as const;
export type ListAccountConnectionsSuspenseQueryKey = ReturnType<
  typeof listAccountConnectionsSuspenseQueryKey
>;
export function listAccountConnectionsSuspenseQueryOptions(
  options: ListAccountConnections["client"]["parameters"] = {},
) {
  const queryKey = listAccountConnectionsSuspenseQueryKey();
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListAccountConnections["data"],
        ListAccountConnections["error"]
      >({
        method: "get",
        url: `/gh/account-connections`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/account-connections
 */
export function useListAccountConnectionsSuspense<
  TData = ListAccountConnections["response"],
  TQueryKey extends QueryKey = ListAccountConnectionsSuspenseQueryKey,
>(
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListAccountConnections["response"],
        ListAccountConnections["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListAccountConnections["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListAccountConnections["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listAccountConnectionsSuspenseQueryKey();
  const query = useSuspenseQuery({
    ...(listAccountConnectionsSuspenseQueryOptions(
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListAccountConnections["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
