import { useQueryClient } from "@tanstack/react-query";
import { FaSync } from "react-icons/fa";
import { useNavigate } from "react-router-dom";

import { dbDeploymentStatus } from "@/api";
import {
  listDeploymentsQueryKey,
  useDeployRepo,
  useListDeployments,
} from "@/api/hooks/deployments";
import { useParams } from "@/api/hooks/router";
import { Button } from "@/components/button";
import { Spinner } from "@/components/icons/spinner";
import { LoadingScreen } from "@/components/loading";
import moment from "moment";
import { useEffect } from "react";
import { useRodelar } from "rodelar/react";
import { StatusIcon } from "../components/status-icon";

export function DeploymentsList() {
  const { id: machineId } = useParams().params as { id: string };
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const rodelar = useRodelar();
  const listDeployments = useListDeployments({ machineId });
  const deployRepo = useDeployRepo({
    mutation: {
      onSuccess: () =>
        queryClient.invalidateQueries({
          queryKey: listDeploymentsQueryKey({ machineId }),
        }),
    },
  });

  useEffect(() => {
    rodelar.subscribe({
      event: `deployment-notification-event/${machineId}`,
      callback() {
        queryClient.invalidateQueries({
          queryKey: listDeploymentsQueryKey({ machineId }),
        });
      },
    });

    // return () => {
    //   rodelar.unsubscribe({
    //     event: `deployment-notification-event/${machineId}`,
    //   });
    // };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (listDeployments.isSuccess && listDeployments.data.length === 0) {
    return (
      <div className="flex-center">
        <label>No Deployments</label>
      </div>
    );
  }

  if (listDeployments.isLoading) {
    return (
      <div className="w-full flex-cneter">
        <LoadingScreen />
      </div>
    );
  }

  return (
    <section className="w-full flex flex-col gap-4">
      <div className="w-full flex-between-center bg-black bg-opacity-20 px-6 py-3 rounded-xl">
        <span className="font-medium block">
          {listDeployments.data?.length} Deployments
        </span>
      </div>

      <div className="w-full border border-gray-600 rounded-xl flex flex-col divide-y divide-gray-600 max-h-[65vh] overflow-scroll">
        {/* TODO: extract views as components */}

        {listDeployments.data?.map((item, index) => (
          <div className="flex justify-between items-center group">
            {item.status != dbDeploymentStatus.DeploymentStatusInProgress && (
              <Button
                type="button"
                className="mx-1 p-3 rounded-full overflow-hidden bg-transparent duration-100 transition-all hidden group-hover:block"
                onClick={() =>
                  deployRepo.mutate({
                    data: {
                      connectionId: item.repoConnectionId,
                      ref: item.commitHash,
                    },
                  })
                }
                disabled={deployRepo.isPending}
              >
                {deployRepo.isPending ? (
                  <Spinner color="white" />
                ) : (
                  <FaSync color="white" />
                )}
              </Button>
            )}

            <div
              key={index}
              onClick={() => navigate(item.id)}
              className="w-full flex justify-between items-center py-3 px-4 cursor-pointer hover:bg-[#333] default-transition"
            >
              <div className="w-[80%] flex gap-4 items-center">
                <span className="w-fit">
                  <StatusIcon status={item.status} />
                </span>
                <span className="font-medium text-ellipsis whitespace-nowrap overflow-hidden">{`${item.commitMessage}`}</span>
              </div>

              <div className="w-[20%] grid">
                <span className="w-full text-sm text-gray-300">
                  {moment(item.createdAt).format("hh:mm DD/MM/yyyy")}
                </span>
                <span className="w-full text-xs text-gray-400">21s</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
}
