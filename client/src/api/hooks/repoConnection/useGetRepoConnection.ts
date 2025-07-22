import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  GetRepoConnectionQueryResponse,
  GetRepoConnectionPathParams,
} from "../../types/GetRepoConnection";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetRepoConnectionClient = typeof client<
  GetRepoConnectionQueryResponse,
  never,
  never
>;
type GetRepoConnection = {
  data: GetRepoConnectionQueryResponse;
  error: never;
  request: never;
  pathParams: GetRepoConnectionPathParams;
  queryParams: never;
  headerParams: never;
  response: GetRepoConnectionQueryResponse;
  client: {
    parameters: Partial<Parameters<GetRepoConnectionClient>[0]>;
    return: Awaited<ReturnType<GetRepoConnectionClient>>;
  };
};
export const getRepoConnectionQueryKey = ({
  connectionId,
}: {
  connectionId: GetRepoConnectionPathParams["connectionId"];
}) =>
  [
    {
      url: "/repo-connection/:connectionId",
      params: { connectionId: connectionId },
    },
  ] as const;
export type GetRepoConnectionQueryKey = ReturnType<
  typeof getRepoConnectionQueryKey
>;
export function getRepoConnectionQueryOptions(
  {
    connectionId,
  }: {
    connectionId: GetRepoConnectionPathParams["connectionId"];
  },
  options: GetRepoConnection["client"]["parameters"] = {},
) {
  const queryKey = getRepoConnectionQueryKey({ connectionId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        GetRepoConnection["data"],
        GetRepoConnection["error"]
      >({
        method: "get",
        url: `/repo-connection/${connectionId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /repo-connection/:connectionId
 */
export function useGetRepoConnection<
  TData = GetRepoConnection["response"],
  TQueryData = GetRepoConnection["response"],
  TQueryKey extends QueryKey = GetRepoConnectionQueryKey,
>(
  {
    connectionId,
  }: {
    connectionId: GetRepoConnectionPathParams["connectionId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetRepoConnection["response"],
        GetRepoConnection["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetRepoConnection["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetRepoConnection["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getRepoConnectionQueryKey({ connectionId });
  const query = useQuery({
    ...(getRepoConnectionQueryOptions(
      { connectionId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetRepoConnection["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getRepoConnectionSuspenseQueryKey = ({
  connectionId,
}: {
  connectionId: GetRepoConnectionPathParams["connectionId"];
}) =>
  [
    {
      url: "/repo-connection/:connectionId",
      params: { connectionId: connectionId },
    },
  ] as const;
export type GetRepoConnectionSuspenseQueryKey = ReturnType<
  typeof getRepoConnectionSuspenseQueryKey
>;
export function getRepoConnectionSuspenseQueryOptions(
  {
    connectionId,
  }: {
    connectionId: GetRepoConnectionPathParams["connectionId"];
  },
  options: GetRepoConnection["client"]["parameters"] = {},
) {
  const queryKey = getRepoConnectionSuspenseQueryKey({ connectionId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        GetRepoConnection["data"],
        GetRepoConnection["error"]
      >({
        method: "get",
        url: `/repo-connection/${connectionId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /repo-connection/:connectionId
 */
export function useGetRepoConnectionSuspense<
  TData = GetRepoConnection["response"],
  TQueryKey extends QueryKey = GetRepoConnectionSuspenseQueryKey,
>(
  {
    connectionId,
  }: {
    connectionId: GetRepoConnectionPathParams["connectionId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetRepoConnection["response"],
        GetRepoConnection["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetRepoConnection["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetRepoConnection["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ??
    getRepoConnectionSuspenseQueryKey({ connectionId });
  const query = useSuspenseQuery({
    ...(getRepoConnectionSuspenseQueryOptions(
      { connectionId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetRepoConnection["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
