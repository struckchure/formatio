import { Controller } from "react-hook-form";

import { DbMachineModel } from "@/api";
import { useCreateRepoConnectionForm } from "@/api/forms/repoConnection";
import { useListRepositories } from "@/api/hooks/github";
import { Button } from "@/components/button";
import { Select } from "@/components/form/select";

export const AddRepoConnectionForm = (props: { machine: DbMachineModel }) => {
  const repoConnectionForm = useCreateRepoConnectionForm({
    machineId: props.machine.id,
  });
  const listRepositories = useListRepositories();

  return (
    <form
      className="flex flex-col gap-4"
      onSubmit={repoConnectionForm.onSubmit}
    >
      <Controller
        name="repoId"
        control={repoConnectionForm.control}
        rules={{
          required: "Required",
          onChange(e) {
            const repoId = e.target.value;
            repoConnectionForm.setValue("repoId", repoId);

            const repoFullname = listRepositories.data?.find(
              (repo) => repo.id === +repoId
            )?.fullName;
            if (repoFullname) {
              repoConnectionForm.setValue("repoName", repoFullname);
            }
          },
        }}
        render={({ field }) => (
          <Select
            label="Repository"
            options={listRepositories.data?.map((repo) => ({
              label: repo.fullName,
              value: repo.id,
            }))}
            error={repoConnectionForm.formState.errors.repoId?.message}
            {...field}
          />
        )}
      />

      <Button
        size="sm"
        className="grid center"
        type="submit"
        isLoading={repoConnectionForm.isPending}
      >
        Connect
      </Button>
    </form>
  );
};
