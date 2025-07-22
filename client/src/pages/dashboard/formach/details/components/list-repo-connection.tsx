import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import toast from "react-hot-toast";
import { AiFillGithub as GithubIcon } from "react-icons/ai";
import { MdClose, MdOpenInNew } from "react-icons/md";
import { Link } from "react-router-dom";

import { DbRepoConnectionModel } from "@/api";
import {
  listRepoConnectionsQueryKey,
  useDeleteRepoConnection,
} from "@/api/hooks/repoConnection";
import { Button } from "@/components/button";
import { Modal } from "@/components/modal";

export function RepoConnectionList(props: {
  connections: DbRepoConnectionModel[];
}) {
  const queryClient = useQueryClient();

  const [deleteConnectionId, setDeleteConnectionId] = useState<string>("");
  const [deleteModalIsVisible, setDeleteModalVisibility] =
    useState<boolean>(false);

  const connectionDeleteMutation = useDeleteRepoConnection({
    mutation: {
      onSuccess: function () {
        toast.success("Repo has been unlinked from machine");

        queryClient.invalidateQueries({
          queryKey: listRepoConnectionsQueryKey(),
        });
      },
    },
  });

  const deleteRepoConnection = (connectionId: string, dryRun = true) => {
    setDeleteConnectionId(connectionId);

    if (dryRun) {
      setDeleteModalVisibility(true);
    } else {
      connectionDeleteMutation.mutate({ connectionId });

      setDeleteModalVisibility(false);
    }
  };

  return (
    <section className="w-full">
      <Modal
        visibility={deleteModalIsVisible}
        setVisibility={setDeleteModalVisibility}
      >
        <div className="text-white flex flex-col gap-4 rounded-lg">
          <p className="font-bold">
            Are you sure you want to unlink this repo from the machine?
          </p>

          <div className="flex items-center justify-end gap-2">
            <Button
              size="sm"
              theme="outline"
              color="danger"
              onClick={() => {
                deleteRepoConnection(deleteConnectionId, false);
              }}
            >
              Yes
            </Button>
            <Button size="sm" onClick={() => setDeleteModalVisibility(false)}>
              No
            </Button>
          </div>
        </div>
      </Modal>

      {props.connections?.map((connection, index) => (
        <div className="flex items-center gap-2 justify-between" key={index}>
          <div className="flex items-center gap-2">
            <label>
              <GithubIcon />
            </label>
            <label>{connection.repoName}</label>
          </div>

          <div className="flex items-center gap-2">
            <Link
              to={`https://github.com/${connection.repoName}`}
              target="_blank"
            >
              <Button className="p-0 m-0" theme="ghost">
                <MdOpenInNew color="white" />
              </Button>
            </Link>

            <Button
              className="p-0 m-0"
              theme="ghost"
              onClick={() => {
                deleteRepoConnection(connection.id);
              }}
            >
              <MdClose color="white" />
            </Button>
          </div>
        </div>
      ))}
    </section>
  );
}
