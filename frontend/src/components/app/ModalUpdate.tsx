import { useEffect, useState } from "react";

//component
import { Button } from "@components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@components/ui/dialog";
import FormCreateCard from "./FormCreateCard";
import { toast, ToastContainer } from "react-toastify";
import Loading from "@components/common/Loading";

//interface
import { UpdateCardPayload } from "@interfaces/card";

//redux
import { useUpdateCardMutation, useGetCardQuery } from "@redux/services/card";
import { ECardType, EVehicleType } from "@constants/enum";

const initForm: UpdateCardPayload = {
  id: "",
  rfid: "",
  owner_name: "",
  card_type: ECardType.MONTH,
  vehicle_type: EVehicleType.MOTORBIKE,
  license_plate: "",
  expired_date: "",
  last_io_history_id: "",
};

interface IProps {
  isOpen: boolean;
  setIsOpen: (value: boolean) => void;
  cardId: string;
}

const ModalUpdate = (props: IProps) => {
  const { isOpen, setIsOpen, cardId } = props;

  const { data: card, isSuccess } = useGetCardQuery(cardId);

  const [formCard, setFormCard] = useState<UpdateCardPayload>(initForm);
  const [UpdateCard, { isLoading }] = useUpdateCardMutation();

  useEffect(() => {
    if (card) {
      setFormCard({
        id: card.id,
        rfid: card.rfid,
        owner_name: card.owner_name,
        card_type: card.card_type,
        vehicle_type: card.vehicle_type,
        license_plate: card.license_plate,
        expired_date: card.expired_date,
        last_io_history_id: card.last_io_history_id
          ? card.last_io_history_id
          : null,
      });
    }
  }, [card]);

  const handleChangeForm = (name: string, value: any) => {
    setFormCard({ ...formCard, [name]: value });
  };

  const handleSubmit = async () => {
    try {
      const result = await UpdateCard(formCard).unwrap();
      if (result) {
        setIsOpen(false);
      }
    } catch (e: any) {
      toast.error(e.data.message, { autoClose: 300 });
    }
  };

  return (
    <>
      {isSuccess && (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Cập nhập Thẻ</DialogTitle>
              <DialogDescription>
                Điền đầy đủ thông tin của thẻ
              </DialogDescription>
            </DialogHeader>
            {formCard.rfid && (
              <FormCreateCard
                handleChangeForm={handleChangeForm}
                formCard={formCard}
              />
            )}
            <DialogFooter>
              <Button disabled={isLoading} type="button" onClick={handleSubmit}>
                {isLoading ? <Loading /> : "Cập nhập"}
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      )}
      <ToastContainer />
    </>
  );
};

export default ModalUpdate;
