import { Controller } from "react-hook-form";

import { CreateNetworkMutationRequest } from "@/api";
import { useCreateNetworkForm } from "@/api/forms/network";
import { useListMachines } from "@/api/hooks/machines";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";
import { Select } from "@/components/form/select";

export function CreateNetworkForm(
  props: {
    initialData?: Partial<CreateNetworkMutationRequest>;
    onSuccess?: () => void;
  } = {}
) {
  const listMachines = useListMachines();
  const createNetworkForm = useCreateNetworkForm(props);

  return (
    <form className="py-10 space-y-6" onSubmit={createNetworkForm.onSubmit}>
      <Controller
        name="machineId"
        control={createNetworkForm.control}
        rules={{ required: "Required" }}
        render={({ field }) => (
          <Select
            className="w-72"
            label="Machine"
            options={
              listMachines.isSuccess
                ? listMachines.data?.map((machine) => ({
                    label: machine.machineName,
                    value: machine.id,
                  }))
                : []
            }
            {...field}
            error={createNetworkForm.formState.errors.machineId?.message}
          />
        )}
      />

      <Controller
        name="destinationPort"
        control={createNetworkForm.control}
        rules={{ required: "Required", min: 1000, max: 9999 }}
        render={({ field }) => (
          <Input
            label="Port"
            type="number"
            className="w-72"
            {...field}
            error={createNetworkForm.formState.errors.destinationPort?.message}
          />
        )}
      />

      <Button isLoading={createNetworkForm.isPending} type="submit">
        Create
      </Button>
    </form>
  );
}
