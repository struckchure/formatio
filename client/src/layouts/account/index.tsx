import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

import AccountHeader from "@/layouts/account/header";
import { useAuthStore } from "@/store";

export function DashboardLayout() {
  const navigate = useNavigate();

  const { isAuthenticated } = useAuthStore();

  useEffect(() => {
    if (!isAuthenticated) {
      navigate(`/login/?next=${window.location.pathname}`);
    }
  }, [isAuthenticated, navigate]);

  return (
    <div className="w-full flex">
      <section className="w-full">
        <AccountHeader />

        <div className="w-full container mx-auto py-10 px-10">
          <Outlet />
        </div>
      </section>
    </div>
  );
}
