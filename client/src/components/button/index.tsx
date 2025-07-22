import * as Headless from "@headlessui/react";
import { ButtonHTMLAttributes, ReactNode } from "react";
import { tv, VariantProps } from "tailwind-variants";

import { popKeys } from "@/utils";

const buttonVariants = tv({
  base: "group text-black hover:bg-opacity-80 rounded-md focus:outline-none data-[focus]:outline-1 duration-100 transition-all cursor-pointer",
  variants: {
    color: {
      primary: "",
      success: "",
      warning: "",
      info: "",
      danger: "text-red-700 border-red-700",
    },
    theme: {
      light: "bg-transparent text-black",
      dark: "bg-white/5 text-white data-[hover]:bg-gray-600 data-[open]:bg-white/5 data-[focus]:outline-white",
      ghost: "bg-transparent p-0",
      outline: "border-[0.5px] bg-transparent",
      link: "border-b-2 p-0 rounded-none",
    },
    size: {
      sm: "px-4 py-2 text-sm",
      md: "px-6 py-3 text-base/6",
      lg: "px-8 py-4 text-lg",
    },
    iconPos: {
      left: "gap-2 flex items-center flex-row",
      right: "gap-2 flex items-center !flex-row-reverse",
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
    iconPos: "left",
  },
});

type ButtonVariants = VariantProps<typeof buttonVariants>;

export interface ButtonProps
  extends Omit<ButtonHTMLAttributes<HTMLButtonElement>, "size"> {
  color?: ButtonVariants["color"];
  theme?: ButtonVariants["theme"];
  size?: ButtonVariants["size"];
  icon?: ReactNode;
  iconPos?: ButtonVariants["iconPos"];
  isLoading?: boolean;
}

export function Button(props: ButtonProps) {
  const { color, theme, size, className, icon } = popKeys(
    props,
    "className",
    "color",
    "theme",
    "size",
    "icon",
    "iconPos",
    "isLoading"
  );

  return (
    <Headless.Button
      {...(props as ButtonHTMLAttributes<HTMLButtonElement>)}
      className={buttonVariants({ color, theme, size, className })}
    >
      <span className="group-hover:cursor-pointer block empty:hidden">
        {icon}
      </span>

      <label className="group-hover:cursor-pointer block empty:hidden">
        {props.children}
      </label>
    </Headless.Button>
  );
}
