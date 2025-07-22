import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListMachinesQueryResponse,
  ListMachinesQueryParams,
} from "../../types/ListMachines";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListMachinesClient = typeof client<
  ListMachinesQueryResponse,
  never,
  never
>;
type ListMachines = {
  data: ListMachinesQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListMachinesQueryParams;
  headerParams: never;
  response: ListMachinesQueryResponse;
  client: {
    parameters: Partial<Parameters<ListMachinesClient>[0]>;
    return: Awaited<ReturnType<ListMachinesClient>>;
  };
};
export const listMachinesQueryKey = (params?: ListMachines["queryParams"]) =>
  [{ url: "/machine" }, ...(params ? [params] : [])] as const;
export type ListMachinesQueryKey = ReturnType<typeof listMachinesQueryKey>;
export function listMachinesQueryOptions(
  params?: ListMachines["queryParams"],
  options: ListMachines["client"]["parameters"] = {},
) {
  const queryKey = listMachinesQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListMachines["data"], ListMachines["error"]>({
        method: "get",
        url: `/machine`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /machine
 */
export function useListMachines<
  TData = ListMachines["response"],
  TQueryData = ListMachines["response"],
  TQueryKey extends QueryKey = ListMachinesQueryKey,
>(
  params?: ListMachines["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListMachines["response"],
        ListMachines["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListMachines["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListMachines["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listMachinesQueryKey(params);
  const query = useQuery({
    ...(listMachinesQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListMachines["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listMachinesSuspenseQueryKey = (
  params?: ListMachines["queryParams"],
) => [{ url: "/machine" }, ...(params ? [params] : [])] as const;
export type ListMachinesSuspenseQueryKey = ReturnType<
  typeof listMachinesSuspenseQueryKey
>;
export function listMachinesSuspenseQueryOptions(
  params?: ListMachines["queryParams"],
  options: ListMachines["client"]["parameters"] = {},
) {
  const queryKey = listMachinesSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListMachines["data"], ListMachines["error"]>({
        method: "get",
        url: `/machine`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /machine
 */
export function useListMachinesSuspense<
  TData = ListMachines["response"],
  TQueryKey extends QueryKey = ListMachinesSuspenseQueryKey,
>(
  params?: ListMachines["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListMachines["response"],
        ListMachines["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListMachines["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListMachines["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? listMachinesSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listMachinesSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListMachines["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
