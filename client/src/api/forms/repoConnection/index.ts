import { useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";

import {
  listRepoConnectionsQueryKey,
  useCreateRepoConnection,
} from "@/api/hooks/repoConnection";
import { CreateRepoConnectionMutationRequest } from "@/api/types";

export function useCreateRepoConnectionForm(
  initialData: Partial<CreateRepoConnectionMutationRequest> = {},
) {
  const queryClient = useQueryClient();

  const mutation = useCreateRepoConnection({
    mutation: {
      onSuccess: function () {
        toast.success("Repo connected successfully");

        queryClient.invalidateQueries({
          queryKey: listRepoConnectionsQueryKey(),
        });
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      },
    },
  });

  const form = useForm<CreateRepoConnectionMutationRequest>();
  const submit: SubmitHandler<CreateRepoConnectionMutationRequest> = (data) =>
    mutation.mutate({ data: { ...initialData, ...data } });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}
