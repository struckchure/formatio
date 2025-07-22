import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListDeploymentsQueryResponse,
  ListDeploymentsQueryParams,
} from "../../types/ListDeployments";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListDeploymentsClient = typeof client<
  ListDeploymentsQueryResponse,
  never,
  never
>;
type ListDeployments = {
  data: ListDeploymentsQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListDeploymentsQueryParams;
  headerParams: never;
  response: ListDeploymentsQueryResponse;
  client: {
    parameters: Partial<Parameters<ListDeploymentsClient>[0]>;
    return: Awaited<ReturnType<ListDeploymentsClient>>;
  };
};
export const listDeploymentsQueryKey = (
  params: ListDeployments["queryParams"],
) => [{ url: "/deployments" }, ...(params ? [params] : [])] as const;
export type ListDeploymentsQueryKey = ReturnType<
  typeof listDeploymentsQueryKey
>;
export function listDeploymentsQueryOptions(
  params: ListDeployments["queryParams"],
  options: ListDeployments["client"]["parameters"] = {},
) {
  const queryKey = listDeploymentsQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListDeployments["data"],
        ListDeployments["error"]
      >({
        method: "get",
        url: `/deployments`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments
 */
export function useListDeployments<
  TData = ListDeployments["response"],
  TQueryData = ListDeployments["response"],
  TQueryKey extends QueryKey = ListDeploymentsQueryKey,
>(
  params: ListDeployments["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListDeployments["response"],
        ListDeployments["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListDeployments["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListDeployments["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listDeploymentsQueryKey(params);
  const query = useQuery({
    ...(listDeploymentsQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListDeployments["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listDeploymentsSuspenseQueryKey = (
  params: ListDeployments["queryParams"],
) => [{ url: "/deployments" }, ...(params ? [params] : [])] as const;
export type ListDeploymentsSuspenseQueryKey = ReturnType<
  typeof listDeploymentsSuspenseQueryKey
>;
export function listDeploymentsSuspenseQueryOptions(
  params: ListDeployments["queryParams"],
  options: ListDeployments["client"]["parameters"] = {},
) {
  const queryKey = listDeploymentsSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListDeployments["data"],
        ListDeployments["error"]
      >({
        method: "get",
        url: `/deployments`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments
 */
export function useListDeploymentsSuspense<
  TData = ListDeployments["response"],
  TQueryKey extends QueryKey = ListDeploymentsSuspenseQueryKey,
>(
  params: ListDeployments["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListDeployments["response"],
        ListDeployments["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListDeployments["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListDeployments["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listDeploymentsSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listDeploymentsSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListDeployments["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
