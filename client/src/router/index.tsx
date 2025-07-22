import {
  BrowserRouter,
  Navigate,
  Outlet,
  Route,
  Routes,
} from "react-router-dom";

import { DashboardLayout } from "@/layouts/account";
import { AuthLayout } from "@/layouts/authlayout";
import { BaseLayout } from "@/layouts/baselayout";
import { HomePage } from "@/pages";
import { ErrorPage } from "@/pages/404";
import { AccountPage } from "@/pages/account";
import { LoginPage } from "@/pages/auth/login";
import { RegisterPage } from "@/pages/auth/register";
import { BillingPage } from "@/pages/billing";
import { Auth0Callback } from "@/pages/callback/auth0";
import { ListFormachPage } from "@/pages/dashboard/formach";
import { CreateFormachPage } from "@/pages/dashboard/formach/create";
import { FormachDetailsPage } from "@/pages/dashboard/formach/details";
import { DeploymentsList } from "@/pages/dashboard/formach/details/deployments";
import { DeploymentLogs } from "@/pages/dashboard/formach/details/deployments/log";
import { MachineOverviewTabPanel } from "@/pages/dashboard/formach/details/machine-overview";
import { NetworkingTabPanel } from "@/pages/dashboard/formach/details/networking";
import { RepositoryTabPanel } from "@/pages/dashboard/formach/details/repository";
import { SettingsTabPanel } from "@/pages/dashboard/formach/details/settings";

export function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="*" element={<ErrorPage />} />

        <Route path="/" element={<AuthLayout />}>
          <Route path="login" element={<LoginPage />} />
          <Route path="register" element={<RegisterPage />} />
        </Route>

        <Route path="/callback/auth0" element={<Auth0Callback />} />

        <Route path="/" element={<BaseLayout />}>
          <Route index element={<HomePage />} />
          <Route path="about" element={<p>about page</p>} />
        </Route>

        <Route path="/dashboard" element={<DashboardLayout />}>
          <Route index element={<Navigate to={"/dashboard/formach"} />} />
          <Route path="formach" element={<ListFormachPage />} />
          <Route path="formach/create" element={<CreateFormachPage />} />
          <Route
            path="formach/:id"
            element={
              <FormachDetailsPage>
                <Outlet />
              </FormachDetailsPage>
            }
          >
            <Route index element={<MachineOverviewTabPanel />} />
            <Route path="repository" element={<RepositoryTabPanel />} />
            <Route path="deployments" element={<DeploymentsList />} />
            <Route
              path="deployments/:deploymentId"
              element={<DeploymentLogs />}
            />
            <Route path="networking" element={<NetworkingTabPanel />} />
            <Route path="settings" element={<SettingsTabPanel />} />
          </Route>

          <Route path="account" element={<AccountPage />} />
          <Route path="billing" element={<BillingPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
