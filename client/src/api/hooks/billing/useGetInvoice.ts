import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  GetInvoiceQueryResponse,
  GetInvoicePathParams,
} from "../../types/GetInvoice";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetInvoiceClient = typeof client<GetInvoiceQueryResponse, never, never>;
type GetInvoice = {
  data: GetInvoiceQueryResponse;
  error: never;
  request: never;
  pathParams: GetInvoicePathParams;
  queryParams: never;
  headerParams: never;
  response: GetInvoiceQueryResponse;
  client: {
    parameters: Partial<Parameters<GetInvoiceClient>[0]>;
    return: Awaited<ReturnType<GetInvoiceClient>>;
  };
};
export const getInvoiceQueryKey = ({
  invoiceId,
}: {
  invoiceId: GetInvoicePathParams["invoiceId"];
}) =>
  [
    { url: "/billing/invoice/:invoiceId", params: { invoiceId: invoiceId } },
  ] as const;
export type GetInvoiceQueryKey = ReturnType<typeof getInvoiceQueryKey>;
export function getInvoiceQueryOptions(
  {
    invoiceId,
  }: {
    invoiceId: GetInvoicePathParams["invoiceId"];
  },
  options: GetInvoice["client"]["parameters"] = {},
) {
  const queryKey = getInvoiceQueryKey({ invoiceId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetInvoice["data"], GetInvoice["error"]>({
        method: "get",
        url: `/billing/invoice/${invoiceId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/invoice/:invoiceId
 */
export function useGetInvoice<
  TData = GetInvoice["response"],
  TQueryData = GetInvoice["response"],
  TQueryKey extends QueryKey = GetInvoiceQueryKey,
>(
  {
    invoiceId,
  }: {
    invoiceId: GetInvoicePathParams["invoiceId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetInvoice["response"],
        GetInvoice["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetInvoice["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetInvoice["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? getInvoiceQueryKey({ invoiceId });
  const query = useQuery({
    ...(getInvoiceQueryOptions(
      { invoiceId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetInvoice["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getInvoiceSuspenseQueryKey = ({
  invoiceId,
}: {
  invoiceId: GetInvoicePathParams["invoiceId"];
}) =>
  [
    { url: "/billing/invoice/:invoiceId", params: { invoiceId: invoiceId } },
  ] as const;
export type GetInvoiceSuspenseQueryKey = ReturnType<
  typeof getInvoiceSuspenseQueryKey
>;
export function getInvoiceSuspenseQueryOptions(
  {
    invoiceId,
  }: {
    invoiceId: GetInvoicePathParams["invoiceId"];
  },
  options: GetInvoice["client"]["parameters"] = {},
) {
  const queryKey = getInvoiceSuspenseQueryKey({ invoiceId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetInvoice["data"], GetInvoice["error"]>({
        method: "get",
        url: `/billing/invoice/${invoiceId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/invoice/:invoiceId
 */
export function useGetInvoiceSuspense<
  TData = GetInvoice["response"],
  TQueryKey extends QueryKey = GetInvoiceSuspenseQueryKey,
>(
  {
    invoiceId,
  }: {
    invoiceId: GetInvoicePathParams["invoiceId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetInvoice["response"],
        GetInvoice["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetInvoice["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetInvoice["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getInvoiceSuspenseQueryKey({ invoiceId });
  const query = useSuspenseQuery({
    ...(getInvoiceSuspenseQueryOptions(
      { invoiceId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetInvoice["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
