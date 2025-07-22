import * as Headless from "@headlessui/react";
import { SelectHTMLAttributes } from "react";
import { tv, VariantProps } from "tailwind-variants";

import { ChevronDownIcon } from "@/components/icons/chevron-down";
import { popKeys } from "@/utils";

const selectVariants = tv({
  base: "block w-full appearance-none rounded-lg border-none focus:outline-none data-[focus]:outline-2 data-[focus]:-outline-offset-2 *:text-black",
  variants: {
    color: {
      light: "bg-transparent text-black",
      dark: "bg-white/5 text-white data-[focus]:outline-white/25",
    },
    size: {
      sm: "text-sm py-1 px-2",
      md: "text-sm/6 py-1.5 px-3",
      lg: "text-lg py-2 px-4",
    },
  },
  compoundVariants: [
    {
      size: ["sm", "md"],
    },
  ],
  defaultVariants: {
    size: "md",
    color: "dark",
  },
});

type SelectVariants = VariantProps<typeof selectVariants>;

interface SelectProps
  extends Omit<SelectHTMLAttributes<HTMLSelectElement>, "color" | "size"> {
  label?: string;
  error?: string;
  description?: string;
  options?: { label: string; value: string | number }[];
  color?: SelectVariants["color"];
  size?: SelectVariants["size"];
}

export function Select(props: SelectProps) {
  const { color, size, className, label, error, description, options } =
    popKeys(
      props,
      "color",
      "size",
      "className",
      "label",
      "error",
      "description",
      "options"
    );

  return (
    <Headless.Field className={className}>
      {label && (
        <Headless.Label className="text-sm/6 font-medium text-white my-2">
          {label}
        </Headless.Label>
      )}

      {description && (
        <Headless.Description className="text-sm/6 text-white/50 my-2">
          {description}
        </Headless.Description>
      )}

      <div className="relative">
        <Headless.Select
          {...(props as SelectHTMLAttributes<HTMLSelectElement>)}
          className={selectVariants({ color, size, className })}
        >
          <option>select {label?.toLowerCase()}</option>
          {options?.map((item, index) => (
            <option key={index} value={item.value}>
              {item.label}
            </option>
          ))}
        </Headless.Select>
        <ChevronDownIcon aria-hidden="true" />
      </div>

      {error && (
        <Headless.Label className="block text-sm/6 font-medium text-red-700 my-2">
          {error}
        </Headless.Label>
      )}
    </Headless.Field>
  );
}
