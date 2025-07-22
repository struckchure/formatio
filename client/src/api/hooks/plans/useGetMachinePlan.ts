import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  GetMachinePlanQueryResponse,
  GetMachinePlanPathParams,
} from "../../types/GetMachinePlan";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetMachinePlanClient = typeof client<
  GetMachinePlanQueryResponse,
  never,
  never
>;
type GetMachinePlan = {
  data: GetMachinePlanQueryResponse;
  error: never;
  request: never;
  pathParams: GetMachinePlanPathParams;
  queryParams: never;
  headerParams: never;
  response: GetMachinePlanQueryResponse;
  client: {
    parameters: Partial<Parameters<GetMachinePlanClient>[0]>;
    return: Awaited<ReturnType<GetMachinePlanClient>>;
  };
};
export const getMachinePlanQueryKey = ({
  machinePlanId,
}: {
  machinePlanId: GetMachinePlanPathParams["machinePlanId"];
}) =>
  [
    { url: "/plans/:machinePlanId", params: { machinePlanId: machinePlanId } },
  ] as const;
export type GetMachinePlanQueryKey = ReturnType<typeof getMachinePlanQueryKey>;
export function getMachinePlanQueryOptions(
  {
    machinePlanId,
  }: {
    machinePlanId: GetMachinePlanPathParams["machinePlanId"];
  },
  options: GetMachinePlan["client"]["parameters"] = {},
) {
  const queryKey = getMachinePlanQueryKey({ machinePlanId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetMachinePlan["data"], GetMachinePlan["error"]>(
        {
          method: "get",
          url: `/plans/${machinePlanId}`,
          ...options,
        },
      );
      return res.data;
    },
  });
}
/**
 * @link /plans/:machinePlanId
 */
export function useGetMachinePlan<
  TData = GetMachinePlan["response"],
  TQueryData = GetMachinePlan["response"],
  TQueryKey extends QueryKey = GetMachinePlanQueryKey,
>(
  {
    machinePlanId,
  }: {
    machinePlanId: GetMachinePlanPathParams["machinePlanId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetMachinePlan["response"],
        GetMachinePlan["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetMachinePlan["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetMachinePlan["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getMachinePlanQueryKey({ machinePlanId });
  const query = useQuery({
    ...(getMachinePlanQueryOptions(
      { machinePlanId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetMachinePlan["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getMachinePlanSuspenseQueryKey = ({
  machinePlanId,
}: {
  machinePlanId: GetMachinePlanPathParams["machinePlanId"];
}) =>
  [
    { url: "/plans/:machinePlanId", params: { machinePlanId: machinePlanId } },
  ] as const;
export type GetMachinePlanSuspenseQueryKey = ReturnType<
  typeof getMachinePlanSuspenseQueryKey
>;
export function getMachinePlanSuspenseQueryOptions(
  {
    machinePlanId,
  }: {
    machinePlanId: GetMachinePlanPathParams["machinePlanId"];
  },
  options: GetMachinePlan["client"]["parameters"] = {},
) {
  const queryKey = getMachinePlanSuspenseQueryKey({ machinePlanId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetMachinePlan["data"], GetMachinePlan["error"]>(
        {
          method: "get",
          url: `/plans/${machinePlanId}`,
          ...options,
        },
      );
      return res.data;
    },
  });
}
/**
 * @link /plans/:machinePlanId
 */
export function useGetMachinePlanSuspense<
  TData = GetMachinePlan["response"],
  TQueryKey extends QueryKey = GetMachinePlanSuspenseQueryKey,
>(
  {
    machinePlanId,
  }: {
    machinePlanId: GetMachinePlanPathParams["machinePlanId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetMachinePlan["response"],
        GetMachinePlan["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetMachinePlan["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetMachinePlan["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getMachinePlanSuspenseQueryKey({ machinePlanId });
  const query = useSuspenseQuery({
    ...(getMachinePlanSuspenseQueryOptions(
      { machinePlanId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetMachinePlan["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
