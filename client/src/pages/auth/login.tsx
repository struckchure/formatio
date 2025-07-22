import { Controller } from "react-hook-form";
import { Link } from "react-router-dom";

import { useLoginUserForm } from "@/api/forms/auth";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";

export function LoginPage() {
  const loginUserForm = useLoginUserForm();

  return (
    <form className="space-y-4 w-full" onSubmit={loginUserForm.onSubmit}>
      <Controller
        name="email"
        control={loginUserForm.control}
        rules={{ required: "Required" }}
        render={({ field }) => (
          <Input
            label="Email"
            className="block w-full"
            type="email"
            {...field}
            error={loginUserForm.formState.errors.email?.message}
          />
        )}
      />

      <Controller
        name="password"
        control={loginUserForm.control}
        rules={{ required: "Required" }}
        render={({ field }) => (
          <Input
            label="Password"
            className="block w-full"
            type="password"
            {...field}
            error={loginUserForm.formState.errors.password?.message}
          />
        )}
      />

      <p className="flex justify-end w-full text-gray-300">
        <Link to="/forgot-password">Forgot Password?</Link>
      </p>

      <div className="flex-center w-full">
        <Button
          className="w-full flex-center"
          isLoading={loginUserForm.isPending}
          type="submit"
        >
          Sign In
        </Button>
      </div>
    </form>
  );
}
