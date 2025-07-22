import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListInvoiceQueryRequest,
  ListInvoiceQueryResponse,
} from "../../types/ListInvoice";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListInvoiceClient = typeof client<
  ListInvoiceQueryResponse,
  never,
  ListInvoiceQueryRequest
>;
type ListInvoice = {
  data: ListInvoiceQueryResponse;
  error: never;
  request: ListInvoiceQueryRequest;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: ListInvoiceQueryResponse;
  client: {
    parameters: Partial<Parameters<ListInvoiceClient>[0]>;
    return: Awaited<ReturnType<ListInvoiceClient>>;
  };
};
export const listInvoiceQueryKey = (data?: ListInvoice["request"]) =>
  [{ url: "/billing/invoice" }, ...(data ? [data] : [])] as const;
export type ListInvoiceQueryKey = ReturnType<typeof listInvoiceQueryKey>;
export function listInvoiceQueryOptions(
  data?: ListInvoice["request"],
  options: ListInvoice["client"]["parameters"] = {},
) {
  const queryKey = listInvoiceQueryKey(data);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListInvoice["data"], ListInvoice["error"]>({
        method: "get",
        url: `/billing/invoice`,
        data,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/invoice
 */
export function useListInvoice<
  TData = ListInvoice["response"],
  TQueryData = ListInvoice["response"],
  TQueryKey extends QueryKey = ListInvoiceQueryKey,
>(
  data?: ListInvoice["request"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListInvoice["response"],
        ListInvoice["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListInvoice["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListInvoice["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listInvoiceQueryKey(data);
  const query = useQuery({
    ...(listInvoiceQueryOptions(
      data,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListInvoice["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listInvoiceSuspenseQueryKey = (data?: ListInvoice["request"]) =>
  [{ url: "/billing/invoice" }, ...(data ? [data] : [])] as const;
export type ListInvoiceSuspenseQueryKey = ReturnType<
  typeof listInvoiceSuspenseQueryKey
>;
export function listInvoiceSuspenseQueryOptions(
  data?: ListInvoice["request"],
  options: ListInvoice["client"]["parameters"] = {},
) {
  const queryKey = listInvoiceSuspenseQueryKey(data);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListInvoice["data"], ListInvoice["error"]>({
        method: "get",
        url: `/billing/invoice`,
        data,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/invoice
 */
export function useListInvoiceSuspense<
  TData = ListInvoice["response"],
  TQueryKey extends QueryKey = ListInvoiceSuspenseQueryKey,
>(
  data?: ListInvoice["request"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListInvoice["response"],
        ListInvoice["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListInvoice["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListInvoice["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listInvoiceSuspenseQueryKey(data);
  const query = useSuspenseQuery({
    ...(listInvoiceSuspenseQueryOptions(
      data,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListInvoice["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
