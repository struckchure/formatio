import { useAuthStore } from "@/store";
import { useNavigate } from "react-router-dom";

import { ACCESS_TOKEN_KEY } from "@/constants/auth";

export function useLogoutUser() {
  const navigate = useNavigate();

  const { setIsAuthenticated } = useAuthStore();

  const logout = () => {
    localStorage.removeItem(ACCESS_TOKEN_KEY);

    setIsAuthenticated(false);

    navigate(`/login/?next=${window.location.pathname}`);
  };

  return { logout };
}
