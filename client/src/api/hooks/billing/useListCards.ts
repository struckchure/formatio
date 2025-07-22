import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  ListCardsQueryResponse,
  ListCardsQueryParams,
} from "../../types/ListCards";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type ListCardsClient = typeof client<ListCardsQueryResponse, never, never>;
type ListCards = {
  data: ListCardsQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: ListCardsQueryParams;
  headerParams: never;
  response: ListCardsQueryResponse;
  client: {
    parameters: Partial<Parameters<ListCardsClient>[0]>;
    return: Awaited<ReturnType<ListCardsClient>>;
  };
};
export const listCardsQueryKey = (params?: ListCards["queryParams"]) =>
  [{ url: "/billing/cards" }, ...(params ? [params] : [])] as const;
export type ListCardsQueryKey = ReturnType<typeof listCardsQueryKey>;
export function listCardsQueryOptions(
  params?: ListCards["queryParams"],
  options: ListCards["client"]["parameters"] = {},
) {
  const queryKey = listCardsQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListCards["data"], ListCards["error"]>({
        method: "get",
        url: `/billing/cards`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/cards
 */
export function useListCards<
  TData = ListCards["response"],
  TQueryData = ListCards["response"],
  TQueryKey extends QueryKey = ListCardsQueryKey,
>(
  params?: ListCards["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        ListCards["response"],
        ListCards["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: ListCards["client"]["parameters"];
  } = {},
): UseQueryResult<TData, ListCards["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listCardsQueryKey(params);
  const query = useQuery({
    ...(listCardsQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, ListCards["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const listCardsSuspenseQueryKey = (params?: ListCards["queryParams"]) =>
  [{ url: "/billing/cards" }, ...(params ? [params] : [])] as const;
export type ListCardsSuspenseQueryKey = ReturnType<
  typeof listCardsSuspenseQueryKey
>;
export function listCardsSuspenseQueryOptions(
  params?: ListCards["queryParams"],
  options: ListCards["client"]["parameters"] = {},
) {
  const queryKey = listCardsSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<ListCards["data"], ListCards["error"]>({
        method: "get",
        url: `/billing/cards`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /billing/cards
 */
export function useListCardsSuspense<
  TData = ListCards["response"],
  TQueryKey extends QueryKey = ListCardsSuspenseQueryKey,
>(
  params?: ListCards["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        ListCards["response"],
        ListCards["error"],
        TData,
        TQueryKey
      >
    >;
    client?: ListCards["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, ListCards["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? listCardsSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(listCardsSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, ListCards["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
