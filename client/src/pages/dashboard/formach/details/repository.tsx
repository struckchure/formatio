import { useGetMachine } from "@/api/hooks/machines";
import { useListRepoConnections } from "@/api/hooks/repoConnection";
import { useParams } from "@/api/hooks/router";
import { GithubAuthorization } from "@/pages/dashboard/formach/details/components/github-authorization";
import { AddRepoConnectionForm } from "./components/add-repo-connection";
import { RepoConnectionList } from "./components/list-repo-connection";

export const RepositoryTabPanel = () => {
  const { id: machineId } = useParams().params as { id: string };

  const getMachine = useGetMachine({ machineId });
  const listRepoConnections = useListRepoConnections({ machineId });

  return (
    <section>
      <div className="flex items-start justify-between gap-6">
        {listRepoConnections.isSuccess &&
        listRepoConnections.data?.length === 0 ? (
          <div className="w-fit mx-auto">
            <div className="w-full flex items-center gap-x-[10rem] justify-between">
              <h3>Connect Repository</h3>

              <GithubAuthorization size="sm" />
            </div>

            {getMachine.isSuccess && (
              <AddRepoConnectionForm machine={getMachine.data} />
            )}
          </div>
        ) : (
          <div className="w-fit">
            {listRepoConnections.isSuccess && (
              <RepoConnectionList connections={listRepoConnections.data} />
            )}
          </div>
        )}
      </div>
    </section>
  );
};
