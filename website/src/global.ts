import create from "zustand";
import { Manga } from "./types";

interface GlobalState {
  publisher: string;
  likedMangas: Manga[];
  changePublisher: (name: string) => void;
  addLikedManga: (manga: Manga) => void;
}

export const useStore = create<GlobalState>((set) => ({
  publisher: "viz",
  likedMangas: [],
  changePublisher: (name: string) => {
    set(() => ({
      publisher: name,
    }));
  },
  addLikedManga: (manga: Manga) => {
    set((state) => ({ likedMangas: [...state.likedMangas, manga] }));
  },
}));
