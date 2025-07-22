import classNames from "classnames";
import { Fragment, ReactNode } from "react";
import { useLocation, useNavigate, useParams } from "react-router-dom";

import { useGetMachine } from "@/api/hooks/machines";
import { VirtualMachineIcon } from "@/components/icons/virtual-machine";
import { LoadingScreen } from "@/components/loading";
import { BackButton } from "@/pages/dashboard/formach/components/back-button";

const tabItems = [
  {
    label: "Overview",
    to: "",
  },
  {
    label: "Repository",
    to: "repository",
  },
  {
    label: "Deployments",
    to: "deployments",
  },
  {
    label: "Networking",
    to: "networking",
  },
  {
    label: "Settings",
    to: "settings",
  },
];

export function FormachDetailsPage(props: { children?: ReactNode }) {
  const params = useParams();
  const navigate = useNavigate();
  const { pathname } = useLocation();

  const activePathname = pathname.split("/")[pathname.split("/").length - 1];

  const getMachine = useGetMachine({ machineId: params.id as string });

  if (getMachine.isLoading && !getMachine.isSuccess) {
    return <LoadingScreen />;
  }

  return (
    <div className="space-y-4">
      <div className="w-full flex justify-start">
        <BackButton />
      </div>
      <Fragment>
        <div className="w-full flex items-start justify-start gap-4">
          <div
            className={classNames(
              "p-4 border-2 rounded-full",
              getMachine.data?.machineStatus == "RUNNING" && "border-success"
            )}
          >
            <VirtualMachineIcon color="#32e0b8" />
          </div>

          <div>
            <h3 className="font-medium text-2xl text-secondary">
              {getMachine.data?.machineName}
            </h3>
            <span className="block text-gray-400">
              cpu: {getMachine.data?.plan.cpu}, memory:{" "}
              {getMachine.data?.plan.memory}
            </span>
          </div>
        </div>

        <div className="w-full flex flex-col">
          <div className="w-full flex justify-start gap-2 top-[4.5rem] py-2 sticky bg-primary">
            {tabItems.map((tabItem, index) => (
              <button
                key={index}
                className={classNames(
                  "py-2 px-4 text-sm font-medium default-transition text-left rounded-lg",
                  activePathname === tabItem.to
                    ? "bg-[#333] text-white"
                    : "text-gray-500 hover:text-white hover:bg-[#555]"
                )}
                onClick={() => navigate(tabItem.to)}
              >
                {tabItem.label}
              </button>
            ))}
          </div>

          <div className="w-full py-10">{props.children}</div>
        </div>
      </Fragment>
    </div>
  );
}
