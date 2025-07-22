import { create } from "zustand";

import { ACCESS_TOKEN_KEY } from "./constants/auth";

type State = {
  isAuthenticated: boolean;
};

type Action = {
  setIsAuthenticated: (lastName: State["isAuthenticated"]) => void;
};

export const useAuthStore = create<State & Action>((set) => ({
  isAuthenticated: localStorage.getItem(ACCESS_TOKEN_KEY) != null,
  setIsAuthenticated: (isAuthenticated: boolean) => set({ isAuthenticated }),
}));
