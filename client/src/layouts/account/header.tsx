import { Link } from "react-router-dom";

import { useLogoutUser } from "@/api/hooks/auth/useLogoutUser";
import { Button } from "@/components/button";
import formatioLogo from "@/images/formatio-logo-48-x-48.png";

function AccountHeader() {
  const { logout } = useLogoutUser();

  return (
    <header className="py-4 px-6 bg-[#333] backdrop-blur-sm w-full sticky top-0">
      <div className="container mx-auto flex justify-between">
        <Link to="/dashboard" className="flex items-center gap-2">
          <img src={formatioLogo} className="object-cover w-10 h-auto" />
          <h1 className="text-lg font-semibold">Formatio</h1>
        </Link>

        <div className="flex items-center justify-end gap-4">
          <Link className="block" to={"/dashboard/formach"}>
            <Button>Machines</Button>
          </Link>

          <Link className="block" to={"/dashboard/account"}>
            <Button>Account</Button>
          </Link>

          <Link className="block" to={"/dashboard/billing"}>
            <Button>Billing</Button>
          </Link>

          <Button onClick={logout} color="danger" theme="outline">
            Logout
          </Button>
        </div>
      </div>
    </header>
  );
}

export default AccountHeader;
