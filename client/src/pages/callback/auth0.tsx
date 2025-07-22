import { useAuth0 } from "@auth0/auth0-react";
import { useEffect, useMemo, useState } from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

import { useAuthSocialConnection } from "@/api/hooks/auth";
import { LoadingScreen } from "@/components/loading";
import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from "@/constants/auth";
import { useAuthStore } from "@/store";

export function Auth0Callback() {
  const navigate = useNavigate();
  const {
    isAuthenticated,
    isLoading: useAuth0IsLoading,
    getAccessTokenSilently,
    getIdTokenClaims,
  } = useAuth0();
  const [token, setToken] = useState<string>("");

  const { setIsAuthenticated } = useAuthStore();

  const socialAuth = useAuthSocialConnection({
    mutation: {
      onSuccess: (data) => {
        localStorage.setItem(ACCESS_TOKEN_KEY, data?.tokens.accessToken);
        localStorage.setItem(REFRESH_TOKEN_KEY, data?.tokens.refreshToken);

        setIsAuthenticated(true);

        navigate("/dashboard");
      },
    },
  });

  const isLoading = useMemo(() => {
    return useAuth0IsLoading || socialAuth.isPending;
  }, [useAuth0IsLoading, socialAuth.isPending]);

  const isError = useMemo(() => {
    return (!isAuthenticated && !isLoading) || socialAuth.isError;
  }, [isAuthenticated, isLoading, socialAuth.isError]);

  useEffect(() => {
    async function getAccessToken() {
      try {
        if (isAuthenticated) {
          const token = await getIdTokenClaims();
          if (token) {
            setToken(token.__raw);
          }
        }
      } catch (error) {
        if (error instanceof Error) toast.error(error.message);
      }
    }

    getAccessToken();
    if (token != "") socialAuth.mutate({ data: { token } });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAuthenticated, getAccessTokenSilently, token]);

  if (isLoading) {
    return <LoadingScreen />;
  }

  if (isError) {
    return (
      <div className="w-full h-screen flex-center">
        <p className="text-red-500 text-lg">
          {JSON.stringify(socialAuth.error) ??
            "Something went wrong, your authentication was not successful"}
        </p>
      </div>
    );
  }

  return (
    <div className="w-full h-screen flex-center">
      <p>Redirecting...</p>
    </div>
  );
}
