import { Fragment, useState } from "react";
import { Controller } from "react-hook-form";

import { useCreateMachineForm } from "@/api/forms/machine";
import { useListMachinePlans } from "@/api/hooks/plans";
import { useGetProfileUser } from "@/api/hooks/user";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";
import { RadioGroup } from "@/components/form/radio";
import { Select } from "@/components/form/select";
import { MACHINE_IMAGE_REPO } from "@/constants/env-vars";
import CreatePlanModal from "@/pages/dashboard/formach/create/components/create-plan-modal";
import { formatCurrency } from "@/utils";
import { FaCheckCircle } from "react-icons/fa";
import { GoPlus } from "react-icons/go";

export function CreateFormachPage() {
  const [createPlanModal, setCreatePlanModal] = useState(false);

  const getProfileUser = useGetProfileUser();
  const listMachinePlans = useListMachinePlans();

  const userIsAdmin =
    getProfileUser.isSuccess && getProfileUser.data?.roles?.includes("ADMIN");

  const createMachineForm = useCreateMachineForm();

  return (
    <Fragment>
      {userIsAdmin && (
        <div className="w-full flex items-center justify-between py-4">
          <CreatePlanModal
            visibility={createPlanModal}
            setVisibility={setCreatePlanModal}
          />

          <Button
            type="button"
            icon={<GoPlus size={20} />}
            onClick={() => setCreatePlanModal(true)}
          >
            Create a plan
          </Button>
        </div>
      )}

      <h1 className="text-3xl font-semibold">Create Formach</h1>
      <form className="py-10 space-x-6 flex items-start">
        <div className="space-y-4">
          <Controller
            name={"machineName"}
            control={createMachineForm.control}
            rules={{ required: "Required" }}
            render={({ field }) => (
              <Input
                label="Machine Name"
                className="w-72"
                {...field}
                error={createMachineForm.formState.errors.machineName?.message}
              />
            )}
          />

          <Controller
            name={"machineImage"}
            control={createMachineForm.control}
            rules={{ required: "Required" }}
            render={({ field }) => (
              <Select
                label="Machine Image"
                description="We currently support Alpine and Ubuntu based images only."
                className="w-72"
                options={[
                  { label: "Alpine", value: `${MACHINE_IMAGE_REPO}/alpine` },
                  { label: "Ubuntu", value: `${MACHINE_IMAGE_REPO}/ubuntu` },
                ]}
                {...field}
                error={createMachineForm.formState.errors.machineImage?.message}
              />
            )}
          />

          <Button
            isLoading={createMachineForm.isPending}
            onClick={createMachineForm.onSubmit}
            className="w-full"
          >
            Create
          </Button>
        </div>

        <RadioGroup
          className="w-full"
          label="Choose a plan"
          options={listMachinePlans.data}
          onChange={(plan) => {
            createMachineForm.setValue("planId", plan.id);
          }}
          render={(plan) => (
            <div className="group relative flex cursor-pointer rounded-lg bg-white/5 py-4 px-5 text-white shadow-md transition focus:outline-none data-[focus]:outline-1 data-[focus]:outline-white data-[checked]:bg-white/10">
              <div className="flex w-full items-center justify-between">
                <div className="text-sm/6">
                  <p className="font-semibold text-white">{plan.name}</p>
                  <div className="flex gap-2 text-white/50">
                    <div>{plan.memory}</div>
                    <div aria-hidden="true">&middot;</div>
                    <div>{plan.cpu}</div>
                    <div aria-hidden="true">&middot;</div>
                    <div>{formatCurrency(plan.monthlyRate)} / month</div>
                  </div>
                </div>

                <FaCheckCircle className="size-6 fill-white opacity-0 transition group-data-[checked]:opacity-100" />
              </div>
            </div>
          )}
        />
      </form>
    </Fragment>
  );
}
