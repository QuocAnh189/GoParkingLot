//component
import { Avatar, AvatarFallback, AvatarImage } from "@components/ui/avatar";

interface IProps {
  avatar_url?: string;
}

const AvatarUser = (props: IProps) => {
  const { avatar_url } = props;
  return (
    <Avatar>
      <AvatarImage
        className="object-cover"
        src={avatar_url ? avatar_url : "https://github.com/shadcn.png"}
        alt="@shadcn"
      />
      <AvatarFallback>CN</AvatarFallback>
    </Avatar>
  );
};

export default AvatarUser;
