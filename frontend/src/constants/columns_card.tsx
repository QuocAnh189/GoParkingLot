//components
import { MoreHorizontal } from "lucide-react";
import { ColumnDef } from "@tanstack/react-table";
import { Checkbox } from "@components/ui/checkbox";
import { Button } from "@components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@components/ui/dropdown-menu";

//interface
import { ICard } from "@interfaces/card";
import { IIOHistory } from "@interfaces/io_history";

//enum
import { ECardType, EVehicleType } from "./enum";

//util
import dayjs from "dayjs";

export const columnsCard = (
  handleDelete: (id: string) => void,
  handleUpdate: (id: string) => void
): ColumnDef<ICard | IIOHistory | any>[] => [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "rfid",
    header: "Mã RFID",
    cell: ({ row }) => <div className="capitalize">{row.getValue("rfid")}</div>,
  },
  {
    accessorKey: "card_type",
    header: "Loại thẻ",
    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue("card_type") === ECardType.MONTH && "Thẻ tháng"}
        {row.getValue("card_type") === ECardType.DAY && "Thẻ ngày"}
      </div>
    ),
  },
  {
    accessorKey: "owner_name",
    header: "Tên chủ thẻ",
    cell: ({ row }) => (
      <div className="capitalize">{row.getValue("owner_name")}</div>
    ),
  },
  {
    accessorKey: "created_at",
    header: "Ngày đăng ký",
    cell: ({ row }) => (
      <div className="capitalize">
        {dayjs(row.getValue("created_at")).format("DD/MM/YYYY")}
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
    accessorKey: "license_plate",
    header: "Biển số xe",
    cell: ({ row }) => (
      <div className="capitalize">{row.getValue("license_plate")}</div>
    ),
  },
  {
    accessorKey: "expired_date",
    header: "Ngày hết hạn",
    cell: ({ row }) => (
      <div className="capitalize">
        {dayjs(row.getValue("expired_date")).format("DD/MM/YYYY")}
      </div>
    ),
  },

  {
    id: "actions",
    enableHiding: false,
    cell: ({ row }) => {
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>

            <DropdownMenuSeparator />
            {/* <DropdownMenuItem>Edit</DropdownMenuItem> */}
            <DropdownMenuItem onClick={() => handleUpdate(row.original.id)}>
              Edit
            </DropdownMenuItem>
            <DropdownMenuItem onClick={() => handleDelete(row.original.id)}>
              Delete
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];
