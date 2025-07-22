import toast from "react-hot-toast";

import {
  useAuthorizeGithubAccount,
  useListAccountConnections,
  useUpdateAppAccess,
} from "@/api/hooks/github";
import { Button, ButtonProps } from "@/components/button";
import { FaGithub } from "react-icons/fa";

export function GithubAuthorization(props: ButtonProps) {
  const authorizeGithub = useAuthorizeGithubAccount(
    { redirectUrl: window.location.href },
    { query: { enabled: false } }
  );
  const updateAppAccess = useUpdateAppAccess({ query: { enabled: false } });
  const listConnections = useListAccountConnections();

  const handleAuthorize = () => {
    authorizeGithub
      .refetch()
      .then((res) => {
        if (res.data) {
          window.location.href = res.data.link;
        }
      })
      .catch((err) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      });
  };

  const handleAppAccessUpdate = () => {
    updateAppAccess
      .refetch()
      .then((res) => {
        if (res.data) {
          window.location.href = res.data.link;
        }
      })
      .catch((err) => {
        toast.error(err?.response?.data?.message ?? err?.message);
      });
  };

  return (
    <div>
      {listConnections.data && listConnections.data.length > 0 ? (
        <Button
          {...props}
          isLoading={updateAppAccess.isLoading || updateAppAccess.isFetching}
          icon={<FaGithub size={20} />}
          onClick={handleAppAccessUpdate}
        >
          Update Github Access
        </Button>
      ) : (
        <Button
          {...props}
          isLoading={authorizeGithub.isLoading || updateAppAccess.isFetching}
          icon={<FaGithub />}
          onClick={handleAuthorize}
        >
          Authorize GitHub
        </Button>
      )}
    </div>
  );
}
