import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListNetworksQueryResponse,
  ListNetworksQueryParams,
} from "../../types/ListNetworks";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListNetworksClient = typeof client<
  ListNetworksQueryResponse,
  never,
  never
>;
type ListNetworks = {
  data: ListNetworksQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListNetworksQueryParams;
  headerParams: never;
  response: ListNetworksQueryResponse;
  client: {
    parameters: Partial<Parameters<ListNetworksClient>[0]>;
    return: Awaited<ReturnType<ListNetworksClient>>;
  };
};
export const listNetworksQueryKey = (params?: ListNetworks["queryParams"]) =>
  [{ url: "/network" }, ...(params ? [params] : [])] as const;
export type ListNetworksQueryKey = ReturnType<typeof listNetworksQueryKey>;
export function listNetworksQueryOptions(
  params?: ListNetworks["queryParams"],
  options: ListNetworks["client"]["parameters"] = {},
) {
  const queryKey = listNetworksQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListNetworks["data"], ListNetworks["error"]>({
        method: "get",
        url: `/network`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /network
 */
export function useListNetworks<
  TData = ListNetworks["response"],
  TQueryData = ListNetworks["response"],
  TQueryKey extends QueryKey = ListNetworksQueryKey,
>(
  params?: ListNetworks["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListNetworks["response"],
        ListNetworks["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListNetworks["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListNetworks["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listNetworksQueryKey(params);
  const query = useQuery({
    ...(listNetworksQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListNetworks["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listNetworksSuspenseQueryKey = (
  params?: ListNetworks["queryParams"],
) => [{ url: "/network" }, ...(params ? [params] : [])] as const;
export type ListNetworksSuspenseQueryKey = ReturnType<
  typeof listNetworksSuspenseQueryKey
>;
export function listNetworksSuspenseQueryOptions(
  params?: ListNetworks["queryParams"],
  options: ListNetworks["client"]["parameters"] = {},
) {
  const queryKey = listNetworksSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListNetworks["data"], ListNetworks["error"]>({
        method: "get",
        url: `/network`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /network
 */
export function useListNetworksSuspense<
  TData = ListNetworks["response"],
  TQueryKey extends QueryKey = ListNetworksSuspenseQueryKey,
>(
  params?: ListNetworks["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListNetworks["response"],
        ListNetworks["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListNetworks["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListNetworks["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listNetworksSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listNetworksSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListNetworks["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
