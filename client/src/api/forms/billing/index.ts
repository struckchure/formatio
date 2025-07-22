import { useQueryClient } from "@tanstack/react-query";
import { SubmitHandler, useForm } from "react-hook-form";

import {
  AuthorizeCardMutationRequest,
  PreAuthorizeCardMutationRequest,
} from "@/api";
import {
  listCardsQueryKey,
  useAuthorizeCard,
  usePreAuthorizeCard,
} from "@/api/hooks/billing";

export function usePreAuthorizeCardForm() {
  const preAuthorizeCard = usePreAuthorizeCard();

  const { register, handleSubmit } = useForm<PreAuthorizeCardMutationRequest>();
  const submit: SubmitHandler<PreAuthorizeCardMutationRequest> = (data) =>
    preAuthorizeCard.mutate({ data });

  return {
    register,
    onSubmit: handleSubmit(submit),
    ...preAuthorizeCard,
  };
}

export function useAuthorizeCardForm(
  initialData: Partial<AuthorizeCardMutationRequest>,
) {
  const queryClient = useQueryClient();
  const authorizeCard = useAuthorizeCard({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: listCardsQueryKey() });
      },
    },
  });

  const { register, handleSubmit } = useForm<AuthorizeCardMutationRequest>();
  const submit: SubmitHandler<AuthorizeCardMutationRequest> = (data) =>
    authorizeCard.mutate({ data: { ...initialData, ...data } });

  return {
    register,
    onSubmit: handleSubmit(submit),
    ...authorizeCard,
  };
}
