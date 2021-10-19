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

import React from "react";
import { useStore } from "../../global";
import { Manga } from "../../types";

interface Props {
  manga: Manga;
}

export const HeartIcon: React.FC<Props> = ({ ...props }) => {
  const { addLikedManga, removeLikedManga, likedMangas } = useStore(
    (state) => state
  );

  const isLikedHelper = () =>
    likedMangas.some((m: Manga) => m.name === props.manga.name);

  const isLiked = (type: string) => {
    if (isLikedHelper()) {
      return "red";
    }
    return type === "stroke" ? "currentColor" : "none";
  };

  const likeManga = () => {
    if (!isLikedHelper()) {
      addLikedManga(props.manga);
    } else {
      removeLikedManga(props.manga);
    }
  };

  return (
    <div onClick={() => likeManga()}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-6 w-6 cursor-pointer"
        fill={isLiked("fill")}
        viewBox="0 0 24 24"
        stroke={isLiked("stroke")}>
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth={2}
          d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
        />
      </svg>
    </div>
  );
};
