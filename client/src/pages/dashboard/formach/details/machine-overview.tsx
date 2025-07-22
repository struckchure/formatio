import { useGetMachine } from "@/api/hooks/machines";
import { useParams } from "@/api/hooks/router";
import { LoadingScreen } from "@/components/loading";

export function MachineOverviewTabPanel() {
  const { id: machineId } = useParams().params as { id: string };
  const getMachine = useGetMachine({ machineId });

  if (getMachine.isLoading) {
    return <LoadingScreen />;
  }

  return (
    <div className="w-full flex items-start justify-start gap-4">
      <div className="w-1/2">
        <h4 className="font-medium text-lg">Status</h4>
        <p className="text-gray-400">{getMachine?.data?.machineStatus}</p>
      </div>

      <div className="w-1/2">
        <h4 className="font-medium text-lg">Image</h4>
        <p className="text-gray-400">{getMachine?.data?.machineImage}</p>
      </div>
    </div>
  );
}
