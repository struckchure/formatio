import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  GetDeploymentByIdQueryResponse,
  GetDeploymentByIdPathParams,
} from "../../types/GetDeploymentById";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetDeploymentByIdClient = typeof client<
  GetDeploymentByIdQueryResponse,
  never,
  never
>;
type GetDeploymentById = {
  data: GetDeploymentByIdQueryResponse;
  error: never;
  request: never;
  pathParams: GetDeploymentByIdPathParams;
  queryParams: never;
  headerParams: never;
  response: GetDeploymentByIdQueryResponse;
  client: {
    parameters: Partial<Parameters<GetDeploymentByIdClient>[0]>;
    return: Awaited<ReturnType<GetDeploymentByIdClient>>;
  };
};
export const getDeploymentByIdQueryKey = ({
  deploymentId,
}: {
  deploymentId: GetDeploymentByIdPathParams["deploymentId"];
}) =>
  [
    {
      url: "/deployments/:deploymentId",
      params: { deploymentId: deploymentId },
    },
  ] as const;
export type GetDeploymentByIdQueryKey = ReturnType<
  typeof getDeploymentByIdQueryKey
>;
export function getDeploymentByIdQueryOptions(
  {
    deploymentId,
  }: {
    deploymentId: GetDeploymentByIdPathParams["deploymentId"];
  },
  options: GetDeploymentById["client"]["parameters"] = {},
) {
  const queryKey = getDeploymentByIdQueryKey({ deploymentId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        GetDeploymentById["data"],
        GetDeploymentById["error"]
      >({
        method: "get",
        url: `/deployments/${deploymentId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments/:deploymentId
 */
export function useGetDeploymentById<
  TData = GetDeploymentById["response"],
  TQueryData = GetDeploymentById["response"],
  TQueryKey extends QueryKey = GetDeploymentByIdQueryKey,
>(
  {
    deploymentId,
  }: {
    deploymentId: GetDeploymentByIdPathParams["deploymentId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetDeploymentById["response"],
        GetDeploymentById["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetDeploymentById["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetDeploymentById["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getDeploymentByIdQueryKey({ deploymentId });
  const query = useQuery({
    ...(getDeploymentByIdQueryOptions(
      { deploymentId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetDeploymentById["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getDeploymentByIdSuspenseQueryKey = ({
  deploymentId,
}: {
  deploymentId: GetDeploymentByIdPathParams["deploymentId"];
}) =>
  [
    {
      url: "/deployments/:deploymentId",
      params: { deploymentId: deploymentId },
    },
  ] as const;
export type GetDeploymentByIdSuspenseQueryKey = ReturnType<
  typeof getDeploymentByIdSuspenseQueryKey
>;
export function getDeploymentByIdSuspenseQueryOptions(
  {
    deploymentId,
  }: {
    deploymentId: GetDeploymentByIdPathParams["deploymentId"];
  },
  options: GetDeploymentById["client"]["parameters"] = {},
) {
  const queryKey = getDeploymentByIdSuspenseQueryKey({ deploymentId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        GetDeploymentById["data"],
        GetDeploymentById["error"]
      >({
        method: "get",
        url: `/deployments/${deploymentId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /deployments/:deploymentId
 */
export function useGetDeploymentByIdSuspense<
  TData = GetDeploymentById["response"],
  TQueryKey extends QueryKey = GetDeploymentByIdSuspenseQueryKey,
>(
  {
    deploymentId,
  }: {
    deploymentId: GetDeploymentByIdPathParams["deploymentId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetDeploymentById["response"],
        GetDeploymentById["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetDeploymentById["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetDeploymentById["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ??
    getDeploymentByIdSuspenseQueryKey({ deploymentId });
  const query = useSuspenseQuery({
    ...(getDeploymentByIdSuspenseQueryOptions(
      { deploymentId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetDeploymentById["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
