import { useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

import { useCreateMachine, useDeleteMachine } from "@/api/hooks/machines";
import { listMachinePlansQueryKey } from "@/api/hooks/plans";
import {
  CreateMachineMutationRequest,
  DeleteMachinePathParams,
} from "@/api/types";

export function useCreateMachineForm(
  initialData: Partial<CreateMachineMutationRequest> = {},
) {
  const navigate = useNavigate();

  const mutation = useCreateMachine({
    mutation: {
      onSuccess: () => {
        toast.success("Machine Creation successful");
        navigate("/dashboard");
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      },
    },
  });

  const form = useForm<CreateMachineMutationRequest>({
    mode: "all",
    defaultValues: initialData,
  });
  const submit: SubmitHandler<CreateMachineMutationRequest> = (data) =>
    mutation.mutate({ data });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}

export function useDeleteMachineForm(initialData: DeleteMachinePathParams) {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const mutation = useDeleteMachine({
    mutation: {
      onSuccess: function () {
        toast.success("Machine deleted successfully");

        queryClient.invalidateQueries({ queryKey: listMachinePlansQueryKey() });
        navigate("/dashboard/formach");
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      },
    },
  });

  const form = useForm<DeleteMachinePathParams>({ defaultValues: initialData });
  const submit: SubmitHandler<DeleteMachinePathParams> = (data) =>
    mutation.mutate(data);

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}
