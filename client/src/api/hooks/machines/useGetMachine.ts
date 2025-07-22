import client from "@/api/client.ts";
import {
  useQuery,
  queryOptions,
  useSuspenseQuery,
} from "@tanstack/react-query";
import type {
  GetMachineQueryResponse,
  GetMachinePathParams,
} from "../../types/GetMachine";
import type {
  QueryObserverOptions,
  UseQueryResult,
  QueryKey,
  UseSuspenseQueryOptions,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";

type GetMachineClient = typeof client<GetMachineQueryResponse, never, never>;
type GetMachine = {
  data: GetMachineQueryResponse;
  error: never;
  request: never;
  pathParams: GetMachinePathParams;
  queryParams: never;
  headerParams: never;
  response: GetMachineQueryResponse;
  client: {
    parameters: Partial<Parameters<GetMachineClient>[0]>;
    return: Awaited<ReturnType<GetMachineClient>>;
  };
};
export const getMachineQueryKey = ({
  machineId,
}: {
  machineId: GetMachinePathParams["machineId"];
}) =>
  [{ url: "/machine/:machineId", params: { machineId: machineId } }] as const;
export type GetMachineQueryKey = ReturnType<typeof getMachineQueryKey>;
export function getMachineQueryOptions(
  {
    machineId,
  }: {
    machineId: GetMachinePathParams["machineId"];
  },
  options: GetMachine["client"]["parameters"] = {},
) {
  const queryKey = getMachineQueryKey({ machineId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetMachine["data"], GetMachine["error"]>({
        method: "get",
        url: `/machine/${machineId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /machine/:machineId
 */
export function useGetMachine<
  TData = GetMachine["response"],
  TQueryData = GetMachine["response"],
  TQueryKey extends QueryKey = GetMachineQueryKey,
>(
  {
    machineId,
  }: {
    machineId: GetMachinePathParams["machineId"];
  },
  options: {
    query?: Partial<
      QueryObserverOptions<
        GetMachine["response"],
        GetMachine["error"],
        TData,
        TQueryData,
        TQueryKey
      >
    >;
    client?: GetMachine["client"]["parameters"];
  } = {},
): UseQueryResult<TData, GetMachine["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey = queryOptions?.queryKey ?? getMachineQueryKey({ machineId });
  const query = useQuery({
    ...(getMachineQueryOptions(
      { machineId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseQueryResult<TData, GetMachine["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
export const getMachineSuspenseQueryKey = ({
  machineId,
}: {
  machineId: GetMachinePathParams["machineId"];
}) =>
  [{ url: "/machine/:machineId", params: { machineId: machineId } }] as const;
export type GetMachineSuspenseQueryKey = ReturnType<
  typeof getMachineSuspenseQueryKey
>;
export function getMachineSuspenseQueryOptions(
  {
    machineId,
  }: {
    machineId: GetMachinePathParams["machineId"];
  },
  options: GetMachine["client"]["parameters"] = {},
) {
  const queryKey = getMachineSuspenseQueryKey({ machineId });
  return queryOptions({
    queryKey,
    queryFn: async () => {
      const res = await client<GetMachine["data"], GetMachine["error"]>({
        method: "get",
        url: `/machine/${machineId}`,
        ...options,
      });
      return res.data;
    },
  });
}
/**
 * @link /machine/:machineId
 */
export function useGetMachineSuspense<
  TData = GetMachine["response"],
  TQueryKey extends QueryKey = GetMachineSuspenseQueryKey,
>(
  {
    machineId,
  }: {
    machineId: GetMachinePathParams["machineId"];
  },
  options: {
    query?: Partial<
      UseSuspenseQueryOptions<
        GetMachine["response"],
        GetMachine["error"],
        TData,
        TQueryKey
      >
    >;
    client?: GetMachine["client"]["parameters"];
  } = {},
): UseSuspenseQueryResult<TData, GetMachine["error"]> & {
  queryKey: TQueryKey;
} {
  const { query: queryOptions, client: clientOptions = {} } = options ?? {};
  const queryKey =
    queryOptions?.queryKey ?? getMachineSuspenseQueryKey({ machineId });
  const query = useSuspenseQuery({
    ...(getMachineSuspenseQueryOptions(
      { machineId },
      clientOptions,
    ) as unknown as QueryObserverOptions),
    queryKey,
    ...(queryOptions as unknown as Omit<QueryObserverOptions, "queryKey">),
  }) as UseSuspenseQueryResult<TData, GetMachine["error"]> & {
    queryKey: TQueryKey;
  };
  query.queryKey = queryKey as TQueryKey;
  return query;
}
