import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  AuthorizeGithubAccountQueryResponse,
  AuthorizeGithubAccountQueryParams,
} from "../../types/AuthorizeGithubAccount";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type AuthorizeGithubAccountClient = typeof client<
  AuthorizeGithubAccountQueryResponse,
  never,
  never
>;
type AuthorizeGithubAccount = {
  data: AuthorizeGithubAccountQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: AuthorizeGithubAccountQueryParams;
  headerParams: never;
  response: AuthorizeGithubAccountQueryResponse;
  client: {
    parameters: Partial<Parameters<AuthorizeGithubAccountClient>[0]>;
    return: Awaited<ReturnType<AuthorizeGithubAccountClient>>;
  };
};
export const authorizeGithubAccountQueryKey = (
  params: AuthorizeGithubAccount["queryParams"],
) => [{ url: "/gh/authorize" }, ...(params ? [params] : [])] as const;
export type AuthorizeGithubAccountQueryKey = ReturnType<
  typeof authorizeGithubAccountQueryKey
>;
export function authorizeGithubAccountQueryOptions(
  params: AuthorizeGithubAccount["queryParams"],
  options: AuthorizeGithubAccount["client"]["parameters"] = {},
) {
  const queryKey = authorizeGithubAccountQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        AuthorizeGithubAccount["data"],
        AuthorizeGithubAccount["error"]
      >({
        method: "get",
        url: `/gh/authorize`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/authorize
 */
export function useAuthorizeGithubAccount<
  TData = AuthorizeGithubAccount["response"],
  TQueryData = AuthorizeGithubAccount["response"],
  TQueryKey extends QueryKey = AuthorizeGithubAccountQueryKey,
>(
  params: AuthorizeGithubAccount["queryParams"],
  options: {
    query?: Partial<
      QueryObserverOptions<
        AuthorizeGithubAccount["response"],
        AuthorizeGithubAccount["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: AuthorizeGithubAccount["client"]["parameters"];
  } = {},
): UseQueryResult<TData, AuthorizeGithubAccount["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? authorizeGithubAccountQueryKey(params);
  const query = useQuery({
    ...(authorizeGithubAccountQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, AuthorizeGithubAccount["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const authorizeGithubAccountSuspenseQueryKey = (
  params: AuthorizeGithubAccount["queryParams"],
) => [{ url: "/gh/authorize" }, ...(params ? [params] : [])] as const;
export type AuthorizeGithubAccountSuspenseQueryKey = ReturnType<
  typeof authorizeGithubAccountSuspenseQueryKey
>;
export function authorizeGithubAccountSuspenseQueryOptions(
  params: AuthorizeGithubAccount["queryParams"],
  options: AuthorizeGithubAccount["client"]["parameters"] = {},
) {
  const queryKey = authorizeGithubAccountSuspenseQueryKey(params);
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<
        AuthorizeGithubAccount["data"],
        AuthorizeGithubAccount["error"]
      >({
        method: "get",
        url: `/gh/authorize`,
        params,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /gh/authorize
 */
export function useAuthorizeGithubAccountSuspense<
  TData = AuthorizeGithubAccount["response"],
  TQueryKey extends QueryKey = AuthorizeGithubAccountSuspenseQueryKey,
>(
  params: AuthorizeGithubAccount["queryParams"],
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        AuthorizeGithubAccount["response"],
        AuthorizeGithubAccount["error"],
        TData,
        TQueryKey
      >
    >;
    client?: AuthorizeGithubAccount["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, AuthorizeGithubAccount["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? authorizeGithubAccountSuspenseQueryKey(params);
  const query = useSuspenseQuery({
    ...(authorizeGithubAccountSuspenseQueryOptions(
      params,
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, AuthorizeGithubAccount["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
