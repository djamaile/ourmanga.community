import create from "zustand";
import produce from "immer";
import { persist } from "zustand/middleware";
import { Manga, Mangas } from "./types";

interface GlobalState {
  publisher: string;
  mangas: Manga[];
  likedMangas: Manga[];
  changePublisher: (name: string) => void;
  addLikedManga: (manga: Manga) => void;
  removeLikedManga: (manga: Manga) => void;
  setMangas: (mangas: Manga[]) => void;
  likeManga: (index: number) => void;
}

export const useStore = create<GlobalState>(
  persist(
    (set, get: any) => ({
      publisher: "viz",
      likedMangas: [],
      mangas: [],
      setMangas: (m: Manga[]) => {
        set(() => ({
          mangas: m,
        }));
      },
      likeManga: (index: number) =>
        set(
          produce((state) => {
            // const manga: Manga = state.mangas.find(
            //   (m: Manga) => m.name === name
            // );
            // manga.liked = !manga.liked;
            console.log(state.mangas);
            const v = !state.mangas[index].liked;
            /* eslint-disable no-param-reassign */
            state.mangas[index].liked = v;
            /* eslint-enable no-param-reassign */
          })
        ),
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
