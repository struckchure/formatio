import { AxiosError } from "axios";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useNavigate, useSearchParams } from "react-router-dom";

import { useLoginUser, useRegisterUser } from "@/api/hooks/auth";
import {
  LoginUserMutationRequest,
  RegisterUserMutationRequest,
} from "@/api/types";
import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from "@/constants/auth";
import { useAuthStore } from "@/store";

export function useLoginUserForm(
  initialData: Partial<LoginUserMutationRequest> = {},
) {
  const navigate = useNavigate();
  const [params] = useSearchParams();

  const { setIsAuthenticated } = useAuthStore();

  const mutation = useLoginUser({
    mutation: {
      onSuccess: (data) => {
        localStorage.setItem(ACCESS_TOKEN_KEY, data.tokens.accessToken);
        localStorage.setItem(REFRESH_TOKEN_KEY, data.tokens.refreshToken);

        toast.success("login successful");

        setIsAuthenticated(true);

        navigate(params.get("next") ?? "/dashboard");
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data.message || err?.message);
      },
    },
  });

  const form = useForm<LoginUserMutationRequest>({
    defaultValues: initialData,
  });
  const submit: SubmitHandler<LoginUserMutationRequest> = (data) =>
    mutation.mutate({ data });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}

export function useRegisterUserForm(
  initialData: Partial<RegisterUserMutationRequest> = {},
) {
  const navigate = useNavigate();

  const mutation = useRegisterUser({
    mutation: {
      onSuccess: (data) => {
        localStorage.setItem(ACCESS_TOKEN_KEY, data.tokens.accessToken);
        localStorage.setItem(REFRESH_TOKEN_KEY, data.tokens.refreshToken);
        toast.success("Registeration successful");
        navigate("/dashboard");
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data?.message || err?.message);
      },
    },
  });

  const form = useForm<RegisterUserMutationRequest>({
    defaultValues: initialData,
  });
  const submit: SubmitHandler<LoginUserMutationRequest> = (data) =>
    mutation.mutate({ data });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}
