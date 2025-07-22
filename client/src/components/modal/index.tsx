import * as Headless from "@headlessui/react";
import { ReactNode } from "react";

export interface ModalProps {
  children?: ReactNode;
  visibility?: boolean;
  setVisibility?: (visibility: boolean) => void;
  showCloseButton?: boolean;
}

export function Modal(props: ModalProps) {
  return (
    <Headless.Dialog
      open={props.visibility}
      as="div"
      className="relative z-10 focus:outline-none"
      onClose={() => props.setVisibility?.(!props.setVisibility)}
    >
      <div className="fixed inset-0 z-10 w-screen overflow-y-auto backdrop-blur-sm duration-100 transition-all">
        <div className="flex min-h-full items-center justify-center p-4">
          <Headless.DialogPanel
            transition
            className="w-full min-w-[30%] max-w-fit rounded-xl bg-white/5 p-6 backdrop-blur-2xl duration-300 ease-out data-[closed]:transform-[scale(95%)] data-[closed]:opacity-0 grid place-items-center"
          >
            {props.children}
          </Headless.DialogPanel>
        </div>
      </div>
    </Headless.Dialog>
  );
}
