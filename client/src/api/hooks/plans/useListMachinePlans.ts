import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListMachinePlansQueryResponse,
  ListMachinePlansQueryParams,
} from "../../types/ListMachinePlans";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListMachinePlansClient = typeof client<
  ListMachinePlansQueryResponse,
  never,
  never
>;
type ListMachinePlans = {
  data: ListMachinePlansQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListMachinePlansQueryParams;
  headerParams: never;
  response: ListMachinePlansQueryResponse;
  client: {
    parameters: Partial<Parameters<ListMachinePlansClient>[0]>;
    return: Awaited<ReturnType<ListMachinePlansClient>>;
  };
};
export const listMachinePlansQueryKey = (
  params?: ListMachinePlans["queryParams"],
) => [{ url: "/plans" }, ...(params ? [params] : [])] as const;
export type ListMachinePlansQueryKey = ReturnType<
  typeof listMachinePlansQueryKey
>;
export function listMachinePlansQueryOptions(
  params?: ListMachinePlans["queryParams"],
  options: ListMachinePlans["client"]["parameters"] = {},
) {
  const queryKey = listMachinePlansQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListMachinePlans["data"],
        ListMachinePlans["error"]
      >({
        method: "get",
        url: `/plans`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /plans
 */
export function useListMachinePlans<
  TData = ListMachinePlans["response"],
  TQueryData = ListMachinePlans["response"],
  TQueryKey extends QueryKey = ListMachinePlansQueryKey,
>(
  params?: ListMachinePlans["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListMachinePlans["response"],
        ListMachinePlans["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListMachinePlans["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListMachinePlans["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listMachinePlansQueryKey(params);
  const query = useQuery({
    ...(listMachinePlansQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListMachinePlans["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listMachinePlansSuspenseQueryKey = (
  params?: ListMachinePlans["queryParams"],
) => [{ url: "/plans" }, ...(params ? [params] : [])] as const;
export type ListMachinePlansSuspenseQueryKey = ReturnType<
  typeof listMachinePlansSuspenseQueryKey
>;
export function listMachinePlansSuspenseQueryOptions(
  params?: ListMachinePlans["queryParams"],
  options: ListMachinePlans["client"]["parameters"] = {},
) {
  const queryKey = listMachinePlansSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        ListMachinePlans["data"],
        ListMachinePlans["error"]
      >({
        method: "get",
        url: `/plans`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /plans
 */
export function useListMachinePlansSuspense<
  TData = ListMachinePlans["response"],
  TQueryKey extends QueryKey = ListMachinePlansSuspenseQueryKey,
>(
  params?: ListMachinePlans["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListMachinePlans["response"],
        ListMachinePlans["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListMachinePlans["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListMachinePlans["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listMachinePlansSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listMachinePlansSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListMachinePlans["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
