import {
  QueryClientProvider as BaseQueryClientProvider,
  QueryClient,
} from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { AxiosError, HttpStatusCode } from "axios";
import toast from "react-hot-toast";

import { ChildrenProps } from "@/types";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retryDelay: 10_000_000,
      retry(failureCount, error) {
        if (error instanceof AxiosError && failureCount >= 3) {
          if (error.status == HttpStatusCode.NotFound) {
            return false;
          }
        }

        return true;
      },
    },
    mutations: {
      onError(error) {
        const message = (
          error as AxiosError<{ message: string; statusCode?: number }>
        ).response?.data?.message;
        if (message) {
          toast.error(message);
        }
      },
    },
  },
});

export function QueryClientProvider({ children }: ChildrenProps) {
  return (
    <BaseQueryClientProvider client={queryClient}>
      {children}
      <ReactQueryDevtools initialIsOpen={false} />
    </BaseQueryClientProvider>
  );
}
