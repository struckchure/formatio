import { useQueryClient } from "@tanstack/react-query";
import { SubmitHandler, useForm } from "react-hook-form";

import { listNetworksQueryKey, useCreateNetwork } from "@/api/hooks/network";
import { CreateNetworkMutationRequest } from "@/api/types";

export function useCreateNetworkForm(props: {
  initialData?: Partial<CreateNetworkMutationRequest>;
  onSuccess?: () => void;
}) {
  const queryClient = useQueryClient();

  const createNetwork = useCreateNetwork({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: listNetworksQueryKey() });

        props.onSuccess?.();
      },
    },
  });

  const form = useForm<CreateNetworkMutationRequest>({
    defaultValues: props.initialData,
  });
  const submit: SubmitHandler<CreateNetworkMutationRequest> = (data) =>
    createNetwork.mutate({
      data: {
        ...data,
        listeningPort: 80,
        destinationPort: +data.destinationPort,
        protocol: "TCP",
      },
    });

  return {
    onSubmit: form.handleSubmit(submit),
    ...form,
    ...createNetwork,
  };
}
