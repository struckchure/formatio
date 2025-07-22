import * as Headless from "@headlessui/react";
import { InputHTMLAttributes } from "react";
import { tv, VariantProps } from "tailwind-variants";

import { popKeys } from "@/utils";

const inputVariants = tv({
  base: "block w-full rounded-lg border-none focus:outline-none data-[focus]:outline-2 data-[focus]:-outline-offset-2",
  variants: {
    theme: {
      light: "bg-transparent text-black",
      dark: "bg-white/5 text-white data-[focus]:outline-white/25",
    },
    size: {
      sm: "px-4 py-2 text-sm",
      md: "px-6 py-3 text-base/6",
      lg: "px-8 py-4 text-lg",
    },
  },
  compoundVariants: [
    {
      size: ["sm", "md"],
    },
  ],
  defaultVariants: {
    size: "sm",
    theme: "dark",
  },
});

type InputVariants = VariantProps<typeof inputVariants>;

interface InputProps
  extends Omit<InputHTMLAttributes<HTMLInputElement>, "size"> {
  theme?: InputVariants["theme"];
  size?: InputVariants["size"];

  label?: string;
  error?: string;
}

export function Input(props: InputProps) {
  const { label, error, theme, size, className } = popKeys(
    props,
    "label",
    "error",
    "theme",
    "size",
    "className"
  );

  return (
    <Headless.Field>
      {label && (
        <Headless.Label className="block text-sm/6 font-medium text-white my-2">
          {label}
        </Headless.Label>
      )}

      <Headless.Input
        {...(props as InputHTMLAttributes<HTMLInputElement>)}
        className={inputVariants({ theme, size, className })}
        autoComplete="off"
      />

      {error && (
        <Headless.Label className="block text-sm/6 font-medium text-red-700 my-2">
          {error}
        </Headless.Label>
      )}
    </Headless.Field>
  );
}
