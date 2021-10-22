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

import { Manga, Mangas } from "../../types";
import { HeartIcon } from "./HeartIcon";

interface Props {
  mangas: Manga[];
}

export const MangaBooks = ({ ...props }: Props) => {
  if (props.mangas === null || props.mangas === undefined) {
    return (
      <div className="grid grid-cols-1">
        <h1 className="text-center capitalize text-4xl">No Manga yet...</h1>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
      {props.mangas.map((manga: Manga) => {
        return (
          <div key={manga.name}>
            <div className="flex flex-col justify-center items-center">
              <div className="bg-yellow-300 w-56 h-72 flex justify-center	items-center rounded-md">
                <img
                  src={manga.image}
                  alt={manga.name}
                  className="w-36 h-56 m-auto block"
                />
              </div>
              <HeartIcon manga={manga} />
              <a
                className="font-bold text-md hover:text-red-500 m-auto block"
                href={manga.link}
                target="_blank"
                rel="noreferrer">
                {manga.name}
              </a>
            </div>
          </div>
        );
      })}
    </div>
  );
};
