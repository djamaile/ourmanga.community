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

import React, { useEffect } from "react";
import { useTitle } from "react-use";
import { useQuery } from "react-query";
import { useStore } from "../../global";
import { PUBLISHERS } from "../../utils/constants";
import { Manga, Mangas } from "../../types";
import { MangaBooks } from "../../components/Home/MangaBooks";
import { PublisherLogo } from "../../components/Home/PublisherLogo";
import { LikedButton } from "../../components/Home/LikedButton";
import mascot from "../../assets/images/mascot.jpg";

const Home: React.FC = () => {
  useTitle("Our Manga - Home");
  const { publisher, likedMangas, setMangas, mangas } = useStore(
    (state) => state
  );
  const date = new Date(Date.now());
  const backend: string =
    process.env.NODE_ENV === "development"
      ? `/releases/${publisher}`
      : `/api/releases/${publisher}`;
  const { data, error, isFetching } = useQuery<Mangas>(["GET", backend, {}]);

  useEffect(() => {
    setMangas(data?.data as Manga[]);
  }, [data]);

  const setLikedMangas = () => {
    setMangas(likedMangas);
  };

  // TODO: make 404 page
  if (error) return <p>${error}</p>;

  return (
    <>
      <div className="container px-4 mx-auto">
        {/* mascot */}
        <div className="grid grid-cols-1 justify-center">
          <section className="flex justify-center">
            <img
              src={mascot}
              alt="mascot"
              width="300"
              height="283"
              className="object-scale-down"
            />
          </section>
        </div>
        <div className="grid grid-cols-1">
          <h1 className="text-4xl black text-center italic mt-4 mb-7">
            <span className="text-red-600">Our Manga</span> for the month of{" "}
            {date.toLocaleString("en-US", { month: "short" })}
          </h1>
        </div>
        {/* publishers */}
        <div className="grid grid-cols-4 lg:grid-cols-5 gap-12 justify-center mb-8">
          {PUBLISHERS.map((p) => {
            return (
              <PublisherLogo
                image={p.image}
                name={p.name}
                key={p.name + p.image}
              />
            );
          })}
        </div>
        {/* Liked Button */}
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
          <div className="flex flex-col justify-center items-center">
            <div className="flex justify-center	items-center">
              <LikedButton setLikedMangas={() => setLikedMangas()} />
            </div>
          </div>
        </div>
        {/* TODO: Seperate to API to MangaBooks, so loading screen doesn't effect the whole screen}
        {/* manga-books */}
        {isFetching ? <></> : <MangaBooks data={mangas} />}
      </div>
    </>
  );
};

export default Home;
