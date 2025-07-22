import * as Headless from "@headlessui/react";
import React from "react";
import { FiChevronRight } from "react-icons/fi";

interface AccordionProps {
  title: React.ReactNode;
  children: React.ReactNode;
}

export function Accordion(props: AccordionProps) {
  return (
    <div className="w-full first:rounded-t-md last:rounded-b-md bg-white/5">
      <Headless.Disclosure as="div" className="px-4 py-3" defaultOpen={false}>
        <Headless.DisclosureButton className="group flex w-full items-center justify-start gap-4">
          <FiChevronRight className="size-5 fill-white/60 group-data-[hover]:fill-white/50 duration-200 transition-all group-data-[open]:rotate-90" />
          <span className="text-sm/6 font-medium text-white group-data-[hover]:text-white/80">
            {props.title}
          </span>
        </Headless.DisclosureButton>
        <Headless.DisclosurePanel>{props.children}</Headless.DisclosurePanel>
      </Headless.Disclosure>
    </div>
  );
}
