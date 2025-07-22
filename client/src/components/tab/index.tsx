import classNames from "classnames";
import { useState } from "react";

import { ChildrenProps } from "@/types";

export interface TabProps {
  labels: string[];
  children: React.ReactNode[];
}

export function Tab(props: TabProps) {
  // TODO: configure vertical and horizontal tab

  const [tabIndex, setTabIndex] = useState(0);

  return (
    <div className="w-full flex flex-col">
      <div className="w-fit flex justify-start gap-2">
        {props.labels.map((label, index) => (
          <button
            className={classNames(
              "py-2 px-4 text-sm font-medium default-transition text-left rounded-lg",
              tabIndex === index
                ? "bg-[#333] text-white"
                : "text-gray-500 hover:text-white hover:bg-[#555]"
            )}
            onClick={() => setTabIndex(index)}
          >
            {label}
          </button>
        ))}
      </div>

      <div className="w-full py-10">
        {props.children[tabIndex]
          ? props.children[tabIndex]
          : "No content found"}
      </div>
    </div>
  );
}

export function TabPanel(props: ChildrenProps) {
  return <div className="w-full">{props.children}</div>;
}
