import { ColumnDef } from "@tanstack/react-table";

//interface
import { ICard } from "@interfaces/card";
import { IIOHistory } from "@interfaces/io_history";

//utils
import dayjs from "dayjs";
import { EVehicleType } from "./enum";

//assets
import noImageDefault from "@assets/noimage.jpg";

export const columnsIO: ColumnDef<ICard | IIOHistory>[] = [
  {
    id: "id",
    accessorKey: "id",
    header: "Mã thẻ",
    cell: ({ row }) => <div className="capitalize">{row.getValue("id")}</div>,
  },
  {
    id: "type",
    accessorKey: "type",
    header: "Trạng thái",
    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue("type") === "IN" ? "Vào" : "Ra"}
      </div>
    ),
  },
  {
    accessorKey: "card_type",
    header: "Loại thẻ",
    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue("card_type") === "MONTH" ? "Thẻ tháng" : "Thẻ ngày"}
      </div>
    ),
  },
  {
    accessorKey: "created_at",
    header: "Giờ vào/ra",
    cell: ({ row }) => (
      <div className="capitalize">
        {dayjs(row.getValue("created_at")).format("DD/MM/YYYY hh:mm:ss A")}
      </div>
    ),
  },
  {
    accessorKey: "vehicle_type",
    header: "Loại xe",
    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue("vehicle_type") === EVehicleType.MOTORBIKE &&
          "Xe gắn máy"}
        {row.getValue("vehicle_type") === EVehicleType.CAR && "Xe ô tô"}
      </div>
    ),
  },
  {
    accessorKey: "crop_url",
    header: "Biển số xe",
    cell: ({ row }) => (
      <div className="capitalize w-20">
        <img
          loading="lazy"
          src={
            row.getValue("crop_url") ? row.getValue("crop_url") : noImageDefault
          }
        />
      </div>
    ),
  },
];
