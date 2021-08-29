import create from "zustand";

interface GlobalState {
  publisher: string;
  changePublisher: (name: string) => void;
}

export const useStore = create<GlobalState>((set) => ({
  publisher: "viz",
  changePublisher: (name: string) => {
    set((state) => ({
      publisher: name,
    }));
  },
}));
