import { useState } from "react";

//component
import { Button } from "@components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@components/ui/dialog";
import FormCreateCard from "./FormCreateCard";
import { toast, ToastContainer } from "react-toastify";
import Loading from "@components/common/Loading";

//icon
import { Plus } from "lucide-react";

//interface
import { CreateCardPayload } from "@interfaces/card";

//redux
import { useCreateCardMutation } from "@redux/services/card";

//enum
import { ECardType, EVehicleType } from "@constants/enum";

const initForm: CreateCardPayload = {
  rfid: "",
  owner_name: "",
  card_type: ECardType.MONTH,
  vehicle_type: EVehicleType.MOTORBIKE,
  license_plate: "",
  expired_date: "",
};

const ModalCreate = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [formCard, setFormCard] = useState<CreateCardPayload>(initForm);
  const [CreateCard, { isLoading }] = useCreateCardMutation();

  const handleChangeForm = (name: string, value: any) => {
    setFormCard({ ...formCard, [name]: value });
  };

  const handleSubmit = async () => {
    try {
      const result = await CreateCard(formCard).unwrap();
      if (result) {
        setFormCard(initForm);
        setIsOpen(false);
      }
    } catch (e: any) {
      toast.error(e.data.message, { autoClose: 300 });
    }
  };

  return (
    <>
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogTrigger asChild>
          <Button
            variant="secondary"
            className="ml-auto hover:cursor-pointer"
            onClick={() => setIsOpen(true)}
          >
            Thêm Thẻ <Plus className="ml-2 h-4 w-4" />
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Thêm Thẻ</DialogTitle>
            <DialogDescription>Điền đầy đủ thông tin của thẻ</DialogDescription>
          </DialogHeader>
          <FormCreateCard
            handleChangeForm={handleChangeForm}
            formCard={formCard}
          />
          <DialogFooter>
            <Button disabled={isLoading} type="button" onClick={handleSubmit}>
              {isLoading ? <Loading /> : "Thêm"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
      <ToastContainer />
    </>
  );
};

export default ModalCreate;
