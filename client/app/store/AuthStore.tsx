import { create } from 'zustand';

export type User = {
  id: string;
  name: string;
  email: string;
  token: string;
};

interface AuthState {
  user: User | null;
  setUser: (user: User) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  setUser: (user) => set({ user }),
}));
