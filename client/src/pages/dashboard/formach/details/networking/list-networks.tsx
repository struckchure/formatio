import { useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useState } from "react";
import toast from "react-hot-toast";

import { DbNetworkModel } from "@/api";
import { listNetworksQueryKey, useDeleteNetwork } from "@/api/hooks/network";
import { Button } from "@/components/button";
import { Modal } from "@/components/modal";

function ListNetworks({ networks }: { networks: DbNetworkModel[] }) {
  const queryClient = useQueryClient();

  const [deleteModalIsVisible, setDeleteModalVisibility] =
    useState<boolean>(false);
  const [deleteNetworkId, setDeleteNetworkId] = useState<string>("");

  const deleteNetworkMutation = useDeleteNetwork({
    mutation: {
      onSuccess: () => {
        toast.success("Network Deleted");

        queryClient.invalidateQueries({ queryKey: listNetworksQueryKey() });
      },
      onError: (err: AxiosError<Error>) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      },
    },
  });

  const deleteNetwork = (networkId: string, dryRun = true) => {
    setDeleteNetworkId(networkId);
    if (dryRun) {
      setDeleteModalVisibility(true);
    } else {
      deleteNetworkMutation.mutate({ networkId });

      setDeleteModalVisibility(false);
    }
  };

  return (
    <section className="w-full">
      <Modal
        visibility={deleteModalIsVisible}
        setVisibility={setDeleteModalVisibility}
      >
        <div className="text-white flex flex-col gap-4">
          <p className="font-bold">
            Are you sure you want to delete this network?
          </p>

          <div className="flex items-center justify-end gap-2">
            <Button
              size="sm"
              theme="outline"
              color="danger"
              onClick={() => {
                deleteNetwork(deleteNetworkId, false);
              }}
            >
              Yes
            </Button>
            <Button
              size="sm"
              onClick={() => setDeleteModalVisibility(!deleteModalIsVisible)}
            >
              No
            </Button>
          </div>
        </div>
      </Modal>

      <div className="flex justify-between">
        <h3 className="font-medium text-lg">Networks</h3>
      </div>

      <div className="mt-4 w-full">
        <table className="w-full">
          <thead>
            <tr className="text-left">
              <th className="font-medium">Machine</th>
              <th className="font-medium">Hostname</th>
              <th className="font-medium">Protocol</th>
              <th className="font-medium">Port</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {networks?.map((network) => (
              <tr key={network.id}>
                <td className="py-2">{network.machine.machineName}</td>
                <td className="py-2">
                  <a
                    href={`http://${network.hostName}`}
                    target="_blank"
                    className="text-blue-300 underline"
                  >
                    {network.hostName}
                  </a>
                </td>
                <td className="py-2">{network.protocol}</td>
                <td className="py-2">{network.destinationPort}</td>
                <td className="w-fit flex gap-2">
                  <Button size="sm">Update</Button>
                  <Button size="sm" onClick={() => deleteNetwork(network.id)}>
                    Delete
                  </Button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </section>
  );
}

export default ListNetworks;
