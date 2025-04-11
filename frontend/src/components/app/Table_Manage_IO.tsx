//hook
import { useState } from "react";

//table
import {
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";

//component
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
} from "@components/ui/dropdown-menu";
import { Label } from "@components/ui/label";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectGroup,
  SelectItem,
} from "@components/ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@components/ui/table";
import Paging from "./Pagination";

//interface
import { columnsIO } from "@constants/columns_io";
import { IListIOHistoryRequest } from "@interfaces/io_history";

//redux
import { useGetIOHistoriesQuery } from "@redux/services/in_out";

//enum
import { ECardType, EVehicleType } from "@constants/enum";
import { Input } from "@components/ui/input";

const initParams: IListIOHistoryRequest = {
  type: "",
  card_type: "",
  vehicle_type: "",
  page: 1,
  size: 8,
  order_by: "",
  order_desc: false,
  take_all: false,
  date: "",
};

const DataTableManageIO = () => {
  const [params, setParams] = useState<IListIOHistoryRequest>(initParams);

  const { data: io_histories } = useGetIOHistoriesQuery(params);

  const table = useReactTable({
    data: io_histories ? io_histories.items : [],
    columns: columnsIO,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    initialState: {
      pagination: {
        pageSize: params.size,
      },
    },
  });

  const handleChangePage = (page: number) => {
    setParams((prevParams) => ({
      ...prevParams,
      page,
    }));
  };

  return (
    <div className="w-full">
      <div className="flex items-center py-4">
        <div className="flex items-center gap-6">
          <div className="">
            <Label htmlFor="name" className="text-right">
              Loại thẻ
            </Label>
            <Select
              name="vehicle_type"
              onValueChange={(value: any) => {
                setParams((prevParams) => ({
                  ...prevParams,
                  card_type: value,
                }));
              }}
            >
              <SelectTrigger className="w-[277px]">
                <SelectValue placeholder="Chọn loại thẻ" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value={ECardType.MONTH}>Thẻ tháng</SelectItem>
                  <SelectItem value={ECardType.DAY}>Thẻ ngày</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
          <div className="">
            <Label htmlFor="name" className="text-right">
              Loại xe
            </Label>
            <Select
              name="vehicle_type"
              onValueChange={(value: any) => {
                setParams((prevParams) => ({
                  ...prevParams,
                  vehicle_type: value,
                }));
              }}
            >
              <SelectTrigger className="w-[277px]">
                <SelectValue placeholder="Chọn loại xe" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value={EVehicleType.MOTORBIKE}>
                    Xe gắn máy
                  </SelectItem>
                  <SelectItem value={EVehicleType.CAR}>Xe ô tô</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>

          <div className="">
            <Label htmlFor="name" className="text-right">
              Ngày
            </Label>
            <Input
              type="date"
              placeholder="Tìm kiếm..."
              className="w-[300px]"
              onChange={(e) => {
                setParams((prevParams) => ({
                  ...prevParams,
                  date: e.target.value,
                }));
              }}
            />
          </div>
        </div>
        <div className="space-x-5 flex justify-end w-full">
          <DropdownMenu>
            <DropdownMenuContent align="end">
              {table
                .getAllColumns()
                .filter((column) => column.getCanHide())
                .map((column) => {
                  return (
                    <DropdownMenuCheckboxItem
                      key={column.id}
                      className="capitalize"
                      checked={column.getIsVisible()}
                      onCheckedChange={(value: any) =>
                        column.toggleVisibility(!!value)
                      }
                    >
                      {column.id}
                    </DropdownMenuCheckboxItem>
                  );
                })}
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                  className={
                    row.getValue("type") === "IN"
                      ? "bg-green-300"
                      : "bg-yellow-400"
                  }
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columnsIO.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="mt-4 flex w-full justify-center">
          {io_histories?.metadata && io_histories.metadata.total_pages > 1 && (
            <Paging
              metadata={io_histories?.metadata}
              onChange={handleChangePage}
            />
          )}
        </div>
      </div>
    </div>
  );
};

export default DataTableManageIO;
