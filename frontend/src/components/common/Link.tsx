//hook
import { useNavigate } from "react-router-dom";

//interface
import { ILink } from "@interfaces/common";

interface IProps {
  link: ILink;
}

const Link = (props: IProps) => {
  const { link } = props;

  const navigate = useNavigate();

  return (
    <button
      onClick={() => {
        navigate(link.path);
      }}
      className="flex items-center gap-2 hover:underline"
    >
      {link.icon}
      <p>{link.name}</p>
    </button>
  );
};

export default Link;
