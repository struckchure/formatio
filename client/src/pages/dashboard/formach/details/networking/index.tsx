import { Suspense, useState } from "react";

import { DbNetworkModel } from "@/api";
import { useListNetworks } from "@/api/hooks/network";
import { useParams } from "@/api/hooks/router";
import { Button } from "@/components/button";
import { LoadingScreen } from "@/components/loading";
import { Modal } from "@/components/modal";
import networkImage from "@/images/network.svg";
import ListNetworks from "@/pages/dashboard/formach/details/networking/list-networks";
import { CreateNetworkForm } from "./create-network-form";

export function NetworkingTabPanel() {
  const { id: machineId } = useParams().params as { id: string };
  const listNetworks = useListNetworks({ machineId: machineId });

  const [createModalIsVisibile, setCreateModalVisibility] =
    useState<boolean>(false);

  return (
    <Suspense fallback={<LoadingScreen />}>
      <Modal
        visibility={createModalIsVisibile}
        setVisibility={setCreateModalVisibility}
      >
        <CreateNetworkForm
          initialData={{ machineId }}
          onSuccess={() => setCreateModalVisibility(!createModalIsVisibile)}
        />
      </Modal>

      {listNetworks.data?.length == 0 ? (
        <section className="flex-col-center gap-2 w-full h-full">
          <div className="bg-[#1e1e1e] p-10 rounded-full">
            <img src={networkImage} alt="virtual machine" className="h-28" />
          </div>
          <h1 className="text-3xl font-medium mt-4">Networking</h1>
          <p className="text-center mt-2">No network configurations yet</p>
          <Button
            onClick={() => setCreateModalVisibility(!createModalIsVisibile)}
          >
            Add Network
          </Button>
        </section>
      ) : (
        <>
          <Button
            size="sm"
            onClick={() => setCreateModalVisibility(!createModalIsVisibile)}
          >
            Create Network
          </Button>
          <ListNetworks networks={listNetworks.data as DbNetworkModel[]} />
        </>
      )}
    </Suspense>
  );
}
