import { useNavigate } from "react-router-dom";

import { Button } from "@/components/button";
import { VirtualMachineIcon } from "@/components/icons/virtual-machine";

function FormachEmptyScreen() {
  const navigate = useNavigate();

  return (
    <section className="flex-col-center gap-2 w-full h-full">
      <div className="bg-[#1e1e1e] p-10 rounded-full">
        <VirtualMachineIcon size={120} color="#32e0b8" />
      </div>
      <h1 className="text-3xl font-medium mt-4">Formachs</h1>
      <p className="font-medium text-lg">Cloud-based virtual machines</p>
      <p className="text-center mt-2">
        Deploy your websites, applications, or any other Cloud-based workloads
        on a scalable and reliable platform.
      </p>
      <Button onClick={() => navigate("/dashboard/formach/create")}>
        Create Formachs
      </Button>
    </section>
  );
}

export default FormachEmptyScreen;
