import * as Headless from "@headlessui/react";
import { Fragment, InputHTMLAttributes, ReactNode } from "react";

export interface RadioInputProps
  extends Omit<InputHTMLAttributes<HTMLInputElement>, "onChange"> {
  checked?: boolean;
  label?: string;
}

export interface RadioGroupProps<T = Record<string, string>>
  extends RadioInputProps {
  options?: T[];
  render: (option: T) => ReactNode;
  onChange?: (value: T) => void;
}

export function RadioGroup<T = unknown>(props: RadioGroupProps<T>) {
  return (
    <Headless.Field>
      {props.label && (
        <Headless.Label className="block text-sm/6 font-medium text-white my-2">
          {props.label}
        </Headless.Label>
      )}

      <Headless.RadioGroup
        value={props.value}
        onChange={(e) => props.onChange?.(e as T)}
        aria-label={props.label}
        className="space-y-2"
      >
        {props.options?.map((option, index) => (
          <Headless.Radio key={index} value={option} as={Fragment}>
            {props.render(option)}
          </Headless.Radio>
        ))}
      </Headless.RadioGroup>
    </Headless.Field>
  );
}
