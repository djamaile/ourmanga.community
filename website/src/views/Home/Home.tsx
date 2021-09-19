import React, { useState, useEffect } from "react";
import { useQuery } from "react-query";
import { useStore } from "../../global";
import { PUBLISHERS } from "../../utils/constants";
import { Manga, Mangas } from "../../types";

interface Props {
  image: string;
  name: string;
}

interface HeartProps {
  manga: Manga;
}

const PublisherLogo: React.FC<Props> = ({ ...props }) => {
  const changePublisher = useStore((state) => state.changePublisher);
  return (
    <section
      className="flex justify-center"
      onClick={() => changePublisher(props.name)}>
      <img
        src={props.image}
        alt="viz media logo"
        width="100"
        height="100"
        className="object-contain cursor-pointer m-auto block"
      />
    </section>
  );
};

const HeartIcon: React.FC<HeartProps> = ({ ...props }) => {
  const [liked, setLiked] = useState<boolean>(false);
  const { addLikedManga, removeLikedManga, likedMangas } = useStore(
    (state) => state
  );

  const isLiked = (type: string) => {
    if (likedMangas.some((m) => m.name === props.manga.name)) {
      return "red";
    }
    return type === "stroke" ? "currentColor" : "none";
  };

  const likeManga = () => {
    if (!liked) {
      addLikedManga(props.manga);
    } else {
      removeLikedManga(props.manga);
    }
    setLiked(!liked);
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

const MangaBooks: React.FC<Mangas> = ({ ...props }) => {
  if (props.data === null || props.data === undefined) {
    return (
      <div className="grid grid-cols-1">
        <h1 className="text-center capitalize text-4xl">No Manga yet...</h1>
      </div>
    );
  }

  return (
    <>
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
        {props.data.map((manga: Manga) => {
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
    </>
  );
};

const Home: React.FC = () => {
  const { publisher, likedMangas, changePublisher } = useStore(
    (state) => state
  );
  const [mangas, setMangas] = React.useState<Manga[]>([]);
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

  if (isFetching) return <p>Is loading...</p>;

  // TODO: make 404 page
  if (error) return <p>${error}</p>;

  return (
    <>
      <div className="container px-4 mx-auto">
        {/* logo */}
        <h1 className="text-4xl black mt-4">私たちの漫画♡</h1>
        {/* publishers */}
        <div className="grid grid-cols-3 lg:grid-cols-5 gap-12 justify-center mb-8">
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
        <div className="grid grid-cols-4 gap-8">
          <div className="flex flex-col justify-center items-center">
            <div className="flex justify-center	items-center">
              <button
                type="button"
                className="bg-yellow-300 hover:bg-yellow-400 text-gray-800 font-bold py-2 px-4 border-yellow-600 hover:border-yellow-500 inline-flex items-center w-56 mb-5"
                onClick={() => setLikedMangas()}>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 cursor-pointer"
                  fill="red"
                  viewBox="0 0 24 24"
                  stroke="red">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={1}
                    d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                  />
                </svg>
                <span>Liked</span>
              </button>
            </div>
          </div>
        </div>
        {/* manga-books */}
        <MangaBooks data={mangas} />
      </div>
    </>
  );
};

export default Home;
