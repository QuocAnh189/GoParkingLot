//hook
import { useEffect, useState } from "react";

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
import { Input } from "@components/ui/input";
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
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
} from "@components/ui/dropdown-menu";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@components/ui/table";
import ModalCreate from "./ModalCreate";
import Paging from "./Pagination";
import ModalUpdate from "./ModalUpdate";
import { toast, ToastContainer } from "react-toastify";

//constant
import { columnsCard } from "@constants/columns_card";
import { useGetCardsQuery, useDeleteCardMutation } from "@redux/services/card";

//interface
import { IListCardRequest } from "@interfaces/card";
import { ECardType, EVehicleType } from "@constants/enum";

//utils
import { useDebounce } from "@utils/useDebounce";

const initParams: IListCardRequest = {
  search: "",
  card_type: "",
  vehicle_type: "",
  page: 1,
  size: 10,
  order_by: "",
  order_desc: false,
  take_all: false,
};

const DataTableManageCard = () => {
  const [isOpenUpdate, setIsOpenUpdate] = useState<boolean>(false);
  const [currentCardId, setCurrentCardId] = useState("");
  const [params, setParams] = useState(initParams);
  const [search, setSearch] = useState<string>("");
  const debouncedSearchTerm = useDebounce(search, 500);

  const { data: cards } = useGetCardsQuery(params);
  const [DeleteCard] = useDeleteCardMutation();

  useEffect(() => {
    setParams({
      ...params,
      search: debouncedSearchTerm,
    });
  }, [debouncedSearchTerm]);

  const handleDelete = async (id: string) => {
    try {
      const result = await DeleteCard(id).unwrap();
      if (result) {
        toast.success("Xóa thẻ thành công", { autoClose: 300 });
      }
    } catch (e: any) {
      toast.error(e.data.message);
    }
  };

  const handleUpdate = (id: string) => {
    setCurrentCardId(id);
    setIsOpenUpdate(true);
  };

  const table = useReactTable({
    data: cards ? cards.items : [],
    columns: columnsCard(handleDelete, handleUpdate),
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
      <div className="flex items-center justify-between py-4 gap-6">
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
        </div>
        <ModalCreate />
        <div className="space-x-5">
          <Input
            placeholder="Tìm kiếm..."
            value={search}
            onChange={(event) => setSearch(event.target.value)}
            className="w-[300px]"
          />
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
            {table.getRowModel().rows?.length &&
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
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
              ))}
          </TableBody>
        </Table>
      </div>
      <div className="mt-4">
        {cards?.metadata && cards.metadata.total_pages > 1 && (
          <Paging metadata={cards?.metadata} onChange={handleChangePage} />
        )}
      </div>
      <ToastContainer autoClose={300} />
      {isOpenUpdate && (
        <ModalUpdate
          isOpen={isOpenUpdate}
          setIsOpen={setIsOpenUpdate}
          cardId={currentCardId}
        />
      )}
    </div>
  );
};

export default DataTableManageCard;
