import { PiWarningLight as WarningIcon } from "react-icons/pi";

import { useGetMachine } from "@/api/hooks/machines";
import { useParams } from "@/api/hooks/router";
import { DeleteMachineForm } from "./components/delete-machine";

export function SettingsTabPanel() {
  const { id: machineId } = useParams().params as { id: string };
  const getMachine = useGetMachine({ machineId });

  return (
    <section className="space-y-4">
      <div>
        <h1 className="text-2xl font-bold">Settings</h1>
        <p className="text-grayText">
          Manage your machine settings and configuration
        </p>
      </div>

      <div className="border-[0.5px] border-danger w-fit rounded-lg p-4 space-y-6">
        <div className="flex items-center gap-2">
          <WarningIcon />
          <h2>Danger Zone</h2>
        </div>

        {getMachine.isSuccess && (
          <DeleteMachineForm machine={getMachine.data} />
        )}
      </div>
    </section>
  );
}
