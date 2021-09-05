import React from "react";
import { useQuery } from "react-query";
import { useStore } from "../../global";
import { PUBLISHERS } from "../../utils/constants";

type Manga = {
  name: string;
  image: string;
  link: string;
};

type Mangas = {
  data: Manga[];
};

interface Props {
  image: string;
  name: string;
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
                    className="m-0 w-36 h-56 m-auto block"
                  />
                </div>
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
  const { data, error, isFetching } = useQuery<Mangas>([
    "GET",
    `/api/releases/${publisher}`,
    {},
  ]);

  const mangas: Manga[] = data?.data as Manga[];

  if (isFetching) return <p>Is loading...</p>;

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
        {/* manga-books */}
        <MangaBooks data={mangas} />
      </div>
    </>
  );
};

export default Home;
