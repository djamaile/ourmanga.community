import { Manga, Mangas } from "../../types";
import { HeartIcon } from "./HeartIcon";

export const MangaBooks: React.FC<Mangas> = ({ ...props }) => {
  if (props.data === null || props.data === undefined) {
    return (
      <div className="grid grid-cols-1">
        <h1 className="text-center capitalize text-4xl">No Manga yet...</h1>
      </div>
    );
  }

  return (
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
  );
};
