import create from "zustand";
import { persist } from "zustand/middleware";
import { Manga } from "./types";

interface GlobalState {
  publisher: string;
  likedMangas: Manga[];
  changePublisher: (name: string) => void;
  addLikedManga: (manga: Manga) => void;
  removeLikedManga: (manga: Manga) => void;
}

export const useStore = create<GlobalState>(
  persist(
    (set, get: any) => ({
      publisher: "viz",
      likedMangas: [],
      changePublisher: (name: string) => {
        set(() => ({
          publisher: name,
        }));
      },
      addLikedManga: (manga: Manga) =>
        set(() => ({ likedMangas: [...get().likedMangas, manga] })),
      removeLikedManga: (manga: Manga) =>
        set(() => ({
          likedMangas: get().likedMangas.filter(
            (m: Manga) => m.name !== manga.name
          ),
        })),
    }),
    {
      name: "mangas", // unique name
    }
  )
);
