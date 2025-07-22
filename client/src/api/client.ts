import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";

import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from "@/constants/auth";
import { API_BASE_URL } from "@/constants/env-vars";

export const axiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: { "Content-Type": "application/json" },
});

// if request status is 401, add token to header
axiosInstance.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${localStorage.getItem(
    ACCESS_TOKEN_KEY,
  )}`;

  return config;
});

// if reponse status is 401, call referesh token api and retry
axiosInstance.interceptors.response.use(
  (response) => response,
  async (error: AxiosError<Error>) => {
    if (
      error.config &&
      error?.response?.status === 401 &&
      !localStorage.getItem(ACCESS_TOKEN_KEY)
    ) {
      // const data = await this.refreshToken();

      // this.storageClass.setItem(ACCESS_TOKEN_KEY, data.access_token);
      // this.storageClass.setItem(REFRESH_TOKEN_KEY, data.refresh_token);

      // this.axiosClient.defaults.headers.common[
      //   "Authorization"
      // ] = `Bearer ${data.accessToken}`;

      // return this.axiosClient(error.config);
      redirectToLogin();
    }
    return Promise.reject(error);
  },
);

function redirectToLogin() {
  window.location.href = `/login/?next=${window.location.pathname}`;

  localStorage.deleteItem(ACCESS_TOKEN_KEY);
  localStorage.deleteItem(REFRESH_TOKEN_KEY);
}

/**
 * Subset of AxiosRequestConfig
 */
type RequestConfig<TData = unknown> = {
  baseURL?: string;
  url?: string;
  method: "get" | "put" | "patch" | "post" | "delete";
  params?: unknown;
  data?: TData;
  responseType?:
    | "arraybuffer"
    | "blob"
    | "document"
    | "json"
    | "text"
    | "stream";
  signal?: AbortSignal;
  headers?: AxiosRequestConfig["headers"];
};

/**
 * Subset of AxiosResponse
 */
type ResponseConfig<TData = unknown> = {
  data: TData;
  status: number;
  statusText: string;
  headers?: AxiosResponse["headers"];
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const axiosClient = async <TData, _TError = unknown, TVariables = unknown>(
  config: RequestConfig<TVariables>,
): Promise<ResponseConfig<TData>> => {
  const promise = axiosInstance.request(config).catch((e) => {
    throw e;
  });

  return promise;
};

export { axiosClient, type RequestConfig, type ResponseConfig };
export default axiosClient;
