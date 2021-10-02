import { useStore } from "../../global";

interface Props {
  image: string;
  name: string;
}

export const PublisherLogo: React.FC<Props> = ({ ...props }) => {
  const changePublisher = useStore((state) => state.changePublisher);
  const publisher = useStore((state) => state.publisher);
  const styles = { divClass: "object-contain cursor-pointer m-auto block" };

  return (
    <section
      className="flex justify-center"
      onClick={() => changePublisher(props.name)}>
      <img
        src={props.image}
        alt={props.name}
        width="100"
        height="100"
        className={`${styles.divClass} 
          ${publisher !== props.name ? "filter grayscale" : ""}`}
      />
    </section>
  );
};
