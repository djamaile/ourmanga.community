interface Props {
  setLikedMangas: () => void;
}

export const LikedButton: React.FC<Props> = ({ ...props }) => {
  return (
    <button
      type="button"
      className="bg-yellow-300 hover:bg-yellow-400 text-gray-800 font-bold py-2 px-4 border-yellow-600 hover:border-yellow-500 inline-flex items-center w-56 mb-5"
      onClick={() => props.setLikedMangas()}>
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
  );
};
