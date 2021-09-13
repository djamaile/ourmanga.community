import React, { useState } from "react";
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
  const addLikedManga = useStore((state) => state.addLikedManga);
  const likedMangas = useStore((state) => state.likedMangas);

  const likeManga = () => {
    if (!liked) {
      addLikedManga(props.manga);
    } else {
      console.log("remove manga");
    }
    setLiked(!liked);
    console.log(likedMangas);
  };

  return (
    <div onClick={() => likeManga()}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-6 w-6 cursor-pointer"
        fill={liked ? "red" : "none"}
        viewBox="0 0 24 24"
        stroke={liked ? "red" : "currentColor"}>
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
  if (props.data === null) {
    return (
      <div className="grid grid-cols-1">
        <h1 className="text-center capitalize text-4xl">
          No releases planned yet...
        </h1>
      </div>
    );
  }

  return (
    <>
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
        {props.data.map((manga: Manga) => {
          return (
            <div>
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
  const publisher = useStore((state) => state.publisher);
  const likedMangas = useStore((state) => state.likedMangas);
  const backend: string =
    process.env.NODE_ENV === "development"
      ? `/releases/${publisher}`
      : `/api/releases/${publisher}`;
  const { data, error, isFetching } = useQuery<Mangas>(["GET", backend, {}]);

  const mangas: Manga[] = data?.data as Manga[];

  if (isFetching) return <p>Is loading...</p>;

  // TODO: make 404 page
  if (error) return <p>${error}</p>;

  console.log(likedMangas);

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
        {/* manga-books */}
        <MangaBooks data={mangas} />
      </div>
    </>
  );
};

export default Home;
