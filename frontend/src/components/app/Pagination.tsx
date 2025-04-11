import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@components/ui/pagination";
import { IMetadata } from "@interfaces/common";

interface IProps {
  metadata: IMetadata;
  onChange: (page: number) => void;
}
const Paging = (props: IProps) => {
  const { metadata, onChange } = props;

  const pages = Array.from({ length: metadata.total_pages }, (_, i) => i + 1);

  return (
    <Pagination>
      <PaginationContent>
        {/* Nút Previous */}
        <PaginationItem>
          {metadata.page !== 1 && (
            <PaginationPrevious onClick={() => onChange(metadata.page - 1)} />
          )}
        </PaginationItem>

        {pages.map((page) => (
          <PaginationItem key={page}>
            <PaginationLink
              onClick={() => onChange(page)}
              isActive={metadata.page === page}
            >
              {page}
            </PaginationLink>
          </PaginationItem>
        ))}

        {/* Nút Next */}
        <PaginationItem>
          {metadata.page !== metadata.total_pages && (
            <PaginationNext onClick={() => onChange(metadata.page + 1)} />
          )}
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
};

export default Paging;
