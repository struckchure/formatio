import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListDeploymentLogsByIdQueryResponse,
  ListDeploymentLogsByIdPathParams,
} from "../../types/ListDeploymentLogsById";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListDeploymentLogsByIdClient = typeof client<
  ListDeploymentLogsByIdQueryResponse,
  never,
  never
>;
type ListDeploymentLogsById = {
  data: ListDeploymentLogsByIdQueryResponse;
  error: never;
  request: never;
  pathParams: ListDeploymentLogsByIdPathParams;
  queryParams: never;
  headerParams: never;
  response: ListDeploymentLogsByIdQueryResponse;
  client: {
    parameters: Partial<Parameters<ListDeploymentLogsByIdClient>[0]>;
    return: Awaited<ReturnType<ListDeploymentLogsByIdClient>>;
  };
};
export const listDeploymentLogsByIdQueryKey = ({
  deploymentId,
}: {
  deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
}) =>
  [
    {
      url: "/deployments/:deploymentId/logs",
      params: { deploymentId: deploymentId },
    },
  ] as const;
export type ListDeploymentLogsByIdQueryKey = ReturnType<
  typeof listDeploymentLogsByIdQueryKey
>;
export function listDeploymentLogsByIdQueryOptions(
  {
    deploymentId,
  }: {
    deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
  },
  options: ListDeploymentLogsById["client"]["parameters"] = {},
) {
  const queryKey = listDeploymentLogsByIdQueryKey({ deploymentId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListDeploymentLogsById["data"],
        ListDeploymentLogsById["error"]
      >({
        method: "get",
        url: `/deployments/${deploymentId}/logs`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments/:deploymentId/logs
 */
export function useListDeploymentLogsById<
  TData = ListDeploymentLogsById["response"],
  TQueryData = ListDeploymentLogsById["response"],
  TQueryKey extends QueryKey = ListDeploymentLogsByIdQueryKey,
>(
  {
    deploymentId,
  }: {
    deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListDeploymentLogsById["response"],
        ListDeploymentLogsById["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListDeploymentLogsById["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListDeploymentLogsById["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listDeploymentLogsByIdQueryKey({ deploymentId });
  const query = useQuery({
    ...(listDeploymentLogsByIdQueryOptions(
      { deploymentId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListDeploymentLogsById["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listDeploymentLogsByIdSuspenseQueryKey = ({
  deploymentId,
}: {
  deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
}) =>
  [
    {
      url: "/deployments/:deploymentId/logs",
      params: { deploymentId: deploymentId },
    },
  ] as const;
export type ListDeploymentLogsByIdSuspenseQueryKey = ReturnType<
  typeof listDeploymentLogsByIdSuspenseQueryKey
>;
export function listDeploymentLogsByIdSuspenseQueryOptions(
  {
    deploymentId,
  }: {
    deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
  },
  options: ListDeploymentLogsById["client"]["parameters"] = {},
) {
  const queryKey = listDeploymentLogsByIdSuspenseQueryKey({ deploymentId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListDeploymentLogsById["data"],
        ListDeploymentLogsById["error"]
      >({
        method: "get",
        url: `/deployments/${deploymentId}/logs`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments/:deploymentId/logs
 */
export function useListDeploymentLogsByIdSuspense<
  TData = ListDeploymentLogsById["response"],
  TQueryKey extends QueryKey = ListDeploymentLogsByIdSuspenseQueryKey,
>(
  {
    deploymentId,
  }: {
    deploymentId: ListDeploymentLogsByIdPathParams["deploymentId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListDeploymentLogsById["response"],
        ListDeploymentLogsById["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListDeploymentLogsById["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListDeploymentLogsById["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ??
    listDeploymentLogsByIdSuspenseQueryKey({ deploymentId });
  const query = useSuspenseQuery({
    ...(listDeploymentLogsByIdSuspenseQueryOptions(
      { deploymentId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListDeploymentLogsById["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
