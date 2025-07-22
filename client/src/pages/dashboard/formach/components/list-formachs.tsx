import classNames from "classnames";
import { Suspense } from "react";
import { useNavigate } from "react-router-dom";

import { DbMachineModel } from "@/api";
import { VirtualMachineIcon } from "@/components/icons/virtual-machine";
import { LoadingScreen } from "@/components/loading";

function formatImageOS(imageName: string) {
  return imageName.split("/")[imageName.split("/").length - 1];
}

export default function ListFormach(props: { machines?: DbMachineModel[] }) {
  const navigate = useNavigate();

  return (
    <Suspense fallback={<LoadingScreen />}>
      <div className="flex justify-start items-center gap-5 flex-wrap">
        {props.machines?.map((machine) => (
          <button
            onClick={() => navigate(`/dashboard/formach/${machine.id}`)}
            className="rounded-lg p-5 flex items-center gap-4 bg-[#333] hover:bg-[#555] group default-transition min-w-[300px] min-h-[100px]"
          >
            <div
              className={classNames(
                "p-2 border-2 rounded-full",
                machine.machineStatus == "RUNNING" && "border-success"
              )}
            >
              <VirtualMachineIcon size={30} />
            </div>

            <div className="flex flex-col items-start justify-start">
              <label className="text-white text-xl group-hover:cursor-[inherit]">
                {machine.machineName}
              </label>

              <small className="text-md">Status: {machine.machineStatus}</small>

              <small className="text-md">
                OS: {formatImageOS(machine.machineImage)}
              </small>

              <small className="text-md">
                memory: {machine.plan.memory}, cpu: {machine.plan.cpu}
              </small>
            </div>
          </button>
        ))}
      </div>
    </Suspense>
  );
}
