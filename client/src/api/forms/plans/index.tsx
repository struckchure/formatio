import { useCreateMachinePlan } from "@/api/hooks/plans";
import { CreateMachinePlanMutationRequest } from "@/api/types";
import { useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";

export function useCreateMachinePlanForm(
  initialData: Partial<CreateMachinePlanMutationRequest> = {},
) {
  const queryClient = useQueryClient();

  const mutation = useCreateMachinePlan({
    mutation: {
      onSuccess: () => {
        toast.success("Plan created successfully");

        queryClient.invalidateQueries({ queryKey: ["plans"] });
      },
      onError: (err: AxiosError<Error>) =>
        toast.error(err.response?.data.message ?? err.message),
    },
  });

  const form = useForm<CreateMachinePlanMutationRequest>({
    defaultValues: initialData,
  });
  const submit: SubmitHandler<CreateMachinePlanMutationRequest> = (data) =>
    mutation.mutate({ data: { ...data, monthlyRate: +data.monthlyRate } });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...mutation,
  };
}
