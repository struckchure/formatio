import { useState } from "react";

import { DbMachineModel } from "@/api";
import { useDeleteMachineForm } from "@/api/forms/machine";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";

export function DeleteMachineForm(props: { machine: DbMachineModel }) {
  const [machineName, setMachineName] = useState<string>("");
  const disabled = props.machine.machineName !== machineName;

  const deleteMachineForm = useDeleteMachineForm({
    machineId: props.machine.id,
  });

  return (
    <form
      className="flex flex-col justify-start gap-4"
      onSubmit={deleteMachineForm.onSubmit}
    >
      <Input
        label="Enter machine name"
        className={
          disabled ? "" : "border-b-danger focus:border-danger text-danger"
        }
        value={machineName}
        onChange={(e) => setMachineName(e.target.value)}
      />

      <Button theme="outline" color="danger" disabled={disabled} type="submit">
        Delete
      </Button>
    </form>
  );
}
