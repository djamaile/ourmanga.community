import React, { useState, useEffect } from "react";
import { useQuery } from "react-query";
import { useStore } from "../../global";

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
        className="object-contain cursor-pointer"
      />
    </section>
  );
};

const Home: React.FC = () => {
  const publisher = useStore((state) => state.publisher);
  const { data, error, isFetching } = useQuery<Mangas>([
    "GET",
    `/releases/${publisher}`,
    {},
  ]);
  // const [manga, setManga] = useState<Manga[] | undefined>(undefined);

  if (isFetching) return <p>Is loading...</p>;

  if (error) return <p>${error}</p>;

  const publishers = [
    {
      image:
        "https://media-exp1.licdn.com/dms/image/C4D0BAQEeo_kogsXllw/company-logo_200_200/0/1519856560642?e=2159024400&v=beta&t=7MOhLsBxLaptYbSWMo9lgqUlHHPTTgyJ3ZgKZIiQw4g",
      name: "viz",
    },
    {
      image: "https://upload.wikimedia.org/wikipedia/en/9/99/Yen_Press.png",
      name: "yenpress",
    },
    {
      image:
        "https://upload.wikimedia.org/wikipedia/en/thumb/f/f8/Dark_Horse_Comics_logo.svg/1200px-Dark_Horse_Comics_logo.svg.png",
      name: "darkhorse",
    },
    { image: "https://logodix.com/logo/1914457.png", name: "kodansha" },
    {
      image:
        "https://pbs.twimg.com/profile_images/875779221414699008/r6prXoN2_400x400.jpg",
      name: "sevenseas",
    },
  ];

  return (
    <>
      <div className="container px-4 mx-auto">
        {/* logo */}
        <h1 className="text-4xl black">私たちの漫画<span>♡</span></h1>
        {/* publishers */}
        <div className="grid grid-cols-6 lg:grid-cols-12 gap-12 justify-center mb-4">
          {publishers.map((p) => {
            return <PublisherLogo image={p.image} name={p.name} />;
          })}
        </div>
        {/* manga-books */}
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
          {data?.data.map((manga: Manga) => {
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
      </div>
    </>
  );
};

export default Home;
