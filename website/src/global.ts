// Copyright 2021 Djamaile Rahamat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import create from "zustand";
import { persist } from "zustand/middleware";
import { Manga } from "./types";

interface GlobalState {
  publisher: string;
  mangas: Manga[];
  likedMangas: Manga[];
  changePublisher: (name: string) => void;
  addLikedManga: (manga: Manga) => void;
  removeLikedManga: (manga: Manga) => void;
  setMangas: (mangas: Manga[]) => void;
}

export const useStore = create<GlobalState>(
  persist(
    (set, get) => ({
      publisher: "viz",
      likedMangas: [],
      mangas: [],
      setMangas: (m: Manga[]) => {
        set(() => ({
          mangas: m,
        }));
      },
      // TODO: use this function when DB is setup
      // likeManga: (index: number) =>
      //   set(
      //     produce((state) => {
      //       // const manga: Manga = state.mangas.find(
      //       //   (m: Manga) => m.name === name
      //       // );
      //       // manga.liked = !manga.liked;
      //       console.log(state.mangas);
      //       const v = !state.mangas[index].liked;
      //       /* eslint-disable no-param-reassign */
      //       state.mangas[index].liked = v;
      //       /* eslint-enable no-param-reassign */
      //     })
      //   ),
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
