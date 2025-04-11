import { IoMdHome } from "react-icons/io";
import { MdManageSearch } from "react-icons/md";

const routes = [
  { name: "Trang chủ", icon: <IoMdHome color="black" />, path: "/main" },
  {
    name: "Danh sách quản lý thẻ",
    icon: <MdManageSearch color="black" />,
    path: "/main/manage-card",
  },
  {
    name: "Lịch sử vào ra",
    icon: <MdManageSearch color="black" />,
    path: "/main/manage-io",
  },
  // {
  //   name: "Cài đặt",
  //   icon: <CiSettings color="black" />,
  //   path: "/main/setting",
  // },
];

export default routes;
