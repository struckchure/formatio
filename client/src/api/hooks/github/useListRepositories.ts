import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListRepositoriesQueryResponse,
  ListRepositoriesQueryParams,
} from "../../types/ListRepositories";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListRepositoriesClient = typeof client<
  ListRepositoriesQueryResponse,
  never,
  never
>;
type ListRepositories = {
  data: ListRepositoriesQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListRepositoriesQueryParams;
  headerParams: never;
  response: ListRepositoriesQueryResponse;
  client: {
    parameters: Partial<Parameters<ListRepositoriesClient>[0]>;
    return: Awaited<ReturnType<ListRepositoriesClient>>;
  };
};
export const listRepositoriesQueryKey = (
  params?: ListRepositories["queryParams"],
) => [{ url: "/gh/repos" }, ...(params ? [params] : [])] as const;
export type ListRepositoriesQueryKey = ReturnType<
  typeof listRepositoriesQueryKey
>;
export function listRepositoriesQueryOptions(
  params?: ListRepositories["queryParams"],
  options: ListRepositories["client"]["parameters"] = {},
) {
  const queryKey = listRepositoriesQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListRepositories["data"],
        ListRepositories["error"]
      >({
        method: "get",
        url: `/gh/repos`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/repos
 */
export function useListRepositories<
  TData = ListRepositories["response"],
  TQueryData = ListRepositories["response"],
  TQueryKey extends QueryKey = ListRepositoriesQueryKey,
>(
  params?: ListRepositories["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListRepositories["response"],
        ListRepositories["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListRepositories["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListRepositories["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listRepositoriesQueryKey(params);
  const query = useQuery({
    ...(listRepositoriesQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListRepositories["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listRepositoriesSuspenseQueryKey = (
  params?: ListRepositories["queryParams"],
) => [{ url: "/gh/repos" }, ...(params ? [params] : [])] as const;
export type ListRepositoriesSuspenseQueryKey = ReturnType<
  typeof listRepositoriesSuspenseQueryKey
>;
export function listRepositoriesSuspenseQueryOptions(
  params?: ListRepositories["queryParams"],
  options: ListRepositories["client"]["parameters"] = {},
) {
  const queryKey = listRepositoriesSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListRepositories["data"],
        ListRepositories["error"]
      >({
        method: "get",
        url: `/gh/repos`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/repos
 */
export function useListRepositoriesSuspense<
  TData = ListRepositories["response"],
  TQueryKey extends QueryKey = ListRepositoriesSuspenseQueryKey,
>(
  params?: ListRepositories["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListRepositories["response"],
        ListRepositories["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListRepositories["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListRepositories["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listRepositoriesSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listRepositoriesSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListRepositories["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
