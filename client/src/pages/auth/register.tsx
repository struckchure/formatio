import { Controller } from "react-hook-form";

import { useRegisterUserForm } from "@/api/forms/auth";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";

export function RegisterPage() {
  const registerUserForm = useRegisterUserForm();

  return (
    <form className="space-y-4 w-full" onSubmit={registerUserForm.onSubmit}>
      <Controller
        name="firstName"
        control={registerUserForm.control}
        render={({ field }) => (
          <Input
            {...field}
            label={"First Name"}
            error={registerUserForm.formState.errors.firstName?.message}
          />
        )}
      />

      <Controller
        name="lastName"
        control={registerUserForm.control}
        render={({ field }) => (
          <Input
            {...field}
            label={"Last Name"}
            error={registerUserForm.formState.errors.lastName?.message}
          />
        )}
      />

      <Controller
        name="email"
        control={registerUserForm.control}
        rules={{ required: "Required" }}
        render={({ field }) => (
          <Input
            {...field}
            label={"Email"}
            type="email"
            error={registerUserForm.formState.errors.email?.message}
          />
        )}
      />

      <Controller
        name="password"
        control={registerUserForm.control}
        rules={{ required: "Required" }}
        render={({ field }) => (
          <Input
            {...field}
            label={"Password"}
            type="password"
            error={registerUserForm.formState.errors.password?.message}
          />
        )}
      />

      <div className="flex-center w-full">
        <Button
          className="w-full flex-center"
          isLoading={registerUserForm.isPending}
          type="submit"
        >
          Sign Up
        </Button>
      </div>
    </form>
  );
}
