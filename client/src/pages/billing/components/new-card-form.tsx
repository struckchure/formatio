import {
  useAuthorizeCardForm,
  usePreAuthorizeCardForm,
} from "@/api/forms/billing";
import { Button } from "@/components/button";

export function NewCardForm() {
  const preAuthorizeCardForm = usePreAuthorizeCardForm();
  const authorizeCardForm = useAuthorizeCardForm({
    reference: preAuthorizeCardForm.data?.reference,
  });

  if (preAuthorizeCardForm.isSuccess) {
    return (
      <form
        onSubmit={authorizeCardForm.onSubmit}
        className="w-[500px] py-10 px-5 flex flex-col gap-4"
      >
        <input
          placeholder="OTP"
          type="number"
          {...authorizeCardForm.register("otp", {
            required: true,
            minLength: 5,
            maxLength: 5,
          })}
        />

        <Button type="submit" isLoading={authorizeCardForm.isPending}>
          Save
        </Button>
      </form>
    );
  }

  return (
    <form
      onSubmit={preAuthorizeCardForm.onSubmit}
      className="w-[500px] py-10 px-5 flex flex-col gap-4"
    >
      <input
        placeholder="Card Number"
        {...preAuthorizeCardForm.register("cardNumber", { required: true })}
      />
      <input
        placeholder="CVV"
        {...preAuthorizeCardForm.register("cvv", { required: true })}
      />
      <input
        placeholder="Expiry Month"
        {...preAuthorizeCardForm.register("expiryMonth", { required: true })}
      />
      <input
        placeholder="Expiry Year"
        {...preAuthorizeCardForm.register("expiryYear", { required: true })}
      />
      <input
        placeholder="PIN"
        {...preAuthorizeCardForm.register("pin", { required: true })}
      />

      <Button type="submit" isLoading={preAuthorizeCardForm.isPending}>
        Save
      </Button>
    </form>
  );
}
