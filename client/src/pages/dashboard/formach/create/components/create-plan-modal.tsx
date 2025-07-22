import { Controller } from "react-hook-form";

import { useCreateMachinePlanForm } from "@/api/forms/plans";
import { Button } from "@/components/button";
import { Input } from "@/components/form/input";
import { Modal, ModalProps } from "@/components/modal";

function CreatePlanModal(props: ModalProps) {
  const createPlanForm = useCreateMachinePlanForm({
    currency: "NGN",
    monthlyRate: 0,
  });

  return (
    <Modal
      visibility={props.visibility}
      setVisibility={props.setVisibility}
      showCloseButton
    >
      <div className="px-10 py-6 space-y-6">
        <h4 className="text-2xl text-left text-white w-fit">Create Plan</h4>
        <form
          className="grid grid-cols-2 gap-4"
          onSubmit={createPlanForm.onSubmit}
        >
          <div className="col-span-2">
            <Controller
              name="name"
              control={createPlanForm.control}
              render={({ field }) => (
                <Input
                  label="Name"
                  {...field}
                  error={createPlanForm.formState.errors.name?.message}
                />
              )}
            />
          </div>

          <Controller
            name="currency"
            control={createPlanForm.control}
            render={({ field }) => (
              <Input
                label="Currency"
                {...field}
                error={createPlanForm.formState.errors.currency?.message}
              />
            )}
          />

          <Controller
            name="monthlyRate"
            control={createPlanForm.control}
            render={({ field }) => (
              <Input
                label="Monthly Rate"
                type="number"
                {...field}
                error={createPlanForm.formState.errors.monthlyRate?.message}
              />
            )}
          />

          <Controller
            name="cpu"
            control={createPlanForm.control}
            render={({ field }) => (
              <Input
                label="CPU"
                {...field}
                error={createPlanForm.formState.errors.cpu?.message}
              />
            )}
          />

          <Controller
            name="memory"
            control={createPlanForm.control}
            render={({ field }) => (
              <Input
                label="Memory"
                {...field}
                error={createPlanForm.formState.errors.memory?.message}
              />
            )}
          />

          <div className="col-span-2">
            <Button
              type="submit"
              className="w-full flex-center"
              isLoading={createPlanForm.isPending}
            >
              Create Plan
            </Button>
          </div>
        </form>
      </div>
    </Modal>
  );
}

export default CreatePlanModal;
