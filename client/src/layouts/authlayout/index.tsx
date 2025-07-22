import { useAuth0 } from "@auth0/auth0-react";
import { useEffect } from "react";
import { FaGithub, FaGoogle } from "react-icons/fa";
import { Link, Outlet, useSearchParams } from "react-router-dom";

import { useRouter } from "@/api/hooks/router";
import { Button } from "@/components/button";
import formatioLogo from "@/images/formatio-logo-160-x-160.png";
import { useAuthStore } from "@/store";

export function AuthLayout() {
  const router = useRouter();
  const [searchParams] = useSearchParams();
  const nextUrl = searchParams.get("next") ?? "/dashboard";

  const { loginWithRedirect } = useAuth0();
  const { isAuthenticated } = useAuthStore();

  useEffect(() => {
    if (isAuthenticated) {
      router.goTo(nextUrl);
    }
  }, [isAuthenticated, nextUrl, router]);

  return (
    <div className="flex center w-full h-screen auth-background">
      <section className="flex-col-center bg-transparent gap-4">
        <Link to={"/"}>
          <img src={formatioLogo} alt="logo" />
        </Link>

        <div className="flex-center gap-2">
          <Button
            type="button"
            className={"px-12"}
            onClick={() => router.goTo("/login")}
          >
            Login
          </Button>

          <Button
            type="button"
            className={"px-12"}
            onClick={() => router.goTo("/register")}
          >
            Register
          </Button>
        </div>

        <Outlet />

        <div className="flex items-center w-full gap-2">
          <div className="empty:block w-[60%] h-[0.1px] my-4 bg-gray-500" />
          <label className="text-center text-md w-full">continue with</label>
          <div className="empty:block w-[60%] h-[0.1px] my-4 bg-gray-500" />
        </div>

        <div className="flex-center gap-4">
          <Button
            onClick={() =>
              loginWithRedirect({
                authorizationParams: { connection: "google-oauth2" },
              })
            }
          >
            <FaGoogle size={20} />
          </Button>

          <Button
            onClick={() =>
              loginWithRedirect({
                authorizationParams: { connection: "github" },
              })
            }
          >
            <FaGithub size={20} />
          </Button>
        </div>
      </section>
    </div>
  );
}
