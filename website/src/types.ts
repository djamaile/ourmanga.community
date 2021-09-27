export type Manga = {
  name: string;
  image: string;
  link: string;
  liked: boolean;
};

export type Mangas = {
  data: Manga[];
};
