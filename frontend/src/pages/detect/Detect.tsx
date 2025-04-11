//hook
import { useEffect, useRef, useState } from "react";

//assets
// import co_hon_mot_bien from "@assets/audio/cohonmotbiensoxe.mp3";
import the_chua_quet_vao from "@assets/audio/thechuaquetvao.mp3";
import the_da_het_han from "@assets/audio/thedahethan.mp3";
import the_khong_hop_le from "@assets/audio/thekhonghople.mp3";
import xedangnamtrongbai from "@assets/audio/xedangnamtrongbai.mp3";
import bien_so_xe_khong_dung from "@assets/audio/biensoxekhongdung.mp3";
import dung_dung_vi_tri from "@assets/audio/vuilongdungdungvitri.mp3";
// import xe_khong_co_trong_bai from "@assets/audio/xekhongcotrongbai.mp3";
import xin_moi_xe_vao from "@assets/audio/xinmoixevao.mp3";
import xin_moi_xe_ra from "@assets/audio/xinmoixera.mp3";

//component
import Webcam from "react-webcam";
import { Tabs, TabsList, TabsTrigger } from "@components/ui/tabs";
import { RFIDReaderInput } from "rfid-reader-input";

//redux
import { useEntranceMutation, useExitMutation } from "@redux/services/in_out";

//util
import { URLtoFile } from "@utils/url_to_file";
import dayjs from "dayjs";

//interface
import { IIOHistory } from "@interfaces/io_history";

//assets
import no_image from "@assets/noimage.jpg";
import { ETypeHistory } from "@constants/enum";

const videoConstraints = {
  facingMode: "user",
};

const Detect = () => {
  const webRef1 = useRef<any>(null);
  const webRef2 = useRef<any>(null);

  const [Entrance] = useEntranceMutation();
  const [Exit] = useExitMutation();

  const [openCardReaderWindow, setOpenCardReaderWindow] =
    useState<boolean>(false);

  const [entranceResult, setEntranceResult] = useState<IIOHistory>();
  const [exitResult, setExitResult] = useState<IIOHistory>();

  const [stateInOut, setStateInOut] = useState<ETypeHistory>(ETypeHistory.IN);

  const [serialCard, setSerialCard] = useState("");

  useEffect(() => {
    if (serialCard) {
      handleCheck(serialCard);
    }
  }, [serialCard]);

  const handleCheck = async (code: string) => {
    try {
      const file: any =
        stateInOut === ETypeHistory.IN
          ? await URLtoFile(webRef1?.current.getScreenshot())
          : await URLtoFile(webRef2?.current.getScreenshot());

      const formData = new FormData();
      formData.append("image", file);
      formData.append("rfid", code.slice(code.length - 10, code.length));
      formData.append("type", stateInOut);

      const result: any =
        stateInOut === ETypeHistory.IN
          ? await Entrance(formData).unwrap()
          : await Exit(formData).unwrap();

      if (result) {
        console.log(result);
        if (stateInOut === ETypeHistory.IN) {
          setEntranceResult(result);
        } else {
          setEntranceResult(result.data_in);
          setExitResult(result.data_out);
        }
        const success =
          stateInOut === ETypeHistory.IN
            ? new Audio(xin_moi_xe_vao)
            : new Audio(xin_moi_xe_ra);
        success.play();
      }
    } catch (e: any) {
      switch (e.data.message) {
        case "same type IN": {
          const err = new Audio(xedangnamtrongbai);
          err.play();
          break;
        }

        case "same type OUT": {
          const err = new Audio(the_chua_quet_vao);
          err.play();
          break;
        }

        case "invalid card": {
          const err = new Audio(the_khong_hop_le);
          err.play();
          break;
        }

        case "expired date": {
          const err = new Audio(the_da_het_han);
          err.play();
          break;
        }

        case "no plate": {
          const err = new Audio(dung_dung_vi_tri);
          err.play();
          break;
        }

        case "license_plate not right": {
          const err = new Audio(bien_so_xe_khong_dung);
          err.play();
          break;
        }

        default:
          break;
      }
    }
  };

  useEffect(() => {
    let timer = undefined;
    if (entranceResult?.id) {
      timer = setTimeout(() => {
        setEntranceResult(undefined);
      }, 5000);
    }
    return () => {
      clearTimeout(timer);
    };
  }, [entranceResult?.id]);

  useEffect(() => {
    let timer = undefined;
    if (exitResult?.id) {
      timer = setTimeout(() => {
        setExitResult(undefined);
      }, 5000);
    }
    return () => {
      clearTimeout(timer);
    };
  }, [exitResult?.id]);

  useEffect(() => {
    if (!openCardReaderWindow) {
      setOpenCardReaderWindow(true);
    }
  }, [openCardReaderWindow]);

  return (
    <div className="w-full max-h-screen flex flex-col gap-4 items-center">
      <div className="absolute w-[300px] bg-red flex items-center justify-center top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] z-[2000] opacity-0">
        <RFIDReaderInput
          textTitle="Cho phép quẹt thẻ"
          textBody="Quẹt thẻ vào khung"
          isOpen={openCardReaderWindow}
          onRequestClose={() => {
            // setSerialCard("");
          }}
          handleCodeCardRFID={setSerialCard}
        />
      </div>

      <div className="w-[90%] flex items-center justify-between">
        <div className="flex items-center space-x-2 ">
          <Tabs
            className="w-full"
            value={stateInOut}
            onValueChange={(value: any) => {
              setStateInOut(value);
              setOpenCardReaderWindow(false);
            }}
          >
            <TabsList className="grid w-full grid-cols-2">
              <TabsTrigger
                className="bg-gray text-white"
                value={ETypeHistory.IN}
              >
                Vào
              </TabsTrigger>
              <TabsTrigger
                className="bg-gray text-white"
                value={ETypeHistory.OUT}
              >
                Ra
              </TabsTrigger>
            </TabsList>
          </Tabs>
        </div>

        {/* <div className="flex items-center gap-2">
          <Label htmlFor="airplane-mode">Cho phép quẹt thẻ</Label>
          <div className="flex items-center space-x-2 z-20">
            <Switch
              checked={openCardReaderWindow}
              onCheckedChange={() => {
                setOpenCardReaderWindow(!openCardReaderWindow);
                // setOpenCardReaderWindow(checked);
              }}
            />
            <Label htmlFor="airplane-mode">
              {openCardReaderWindow ? "On" : "Off"}
            </Label>
          </div>
        </div> */}
      </div>

      <div className="w-full flex gap-10 items-center justify-center">
        <div className="w-2/5">
          {entranceResult?.image_url ? (
            <img src={entranceResult?.image_url} className="w-full" />
          ) : (
            <Webcam
              mirrored={true}
              className="w-full"
              ref={webRef1}
              audio={false}
              videoConstraints={videoConstraints}
              onUserMediaError={(err) => console.error("Webcam Error: ", err)}
            />
          )}
        </div>
        <div className="w-2/5">
          {exitResult?.image_url ? (
            <img src={exitResult?.image_url} className="w-full" />
          ) : (
            <Webcam
              mirrored={true}
              className="w-full"
              ref={webRef2}
              audio={false}
              videoConstraints={videoConstraints}
              onUserMediaError={(err) => console.error("Webcam Error: ", err)}
            />
          )}
        </div>
      </div>

      <div className="w-[90%] flex items-center gap-2 shadow-xl rounded-sm mt-2">
        <div className="w-1/2 h-full flex items-center gap-2">
          <div className="w-1/2 h-full">
            <div className="w-full flex items-center p-2 justify-between border-[2px] border-gray rounded-t-sm">
              <p className="text-xl">Ngày/giờ vào:</p>
              <p className="text-xl font-bold">
                {entranceResult?.created_at
                  ? dayjs(entranceResult?.created_at).format(
                      "DD/MM/YYYY hh:mm:ss"
                    )
                  : ""}
              </p>
            </div>
            <div className="w-full shadow-xl py-2 px-2 border-[3px] border-gray rounded-b-sm">
              <p className="text-xl">Biển số xe vào</p>
              <h1 className="text-3xl font-bold text-center">
                <img
                  className={`w-full h-[100px] ${
                    entranceResult?.crop_url ? "object-contain" : "object-cover"
                  }`}
                  src={
                    entranceResult?.crop_url
                      ? entranceResult.crop_url
                      : no_image
                  }
                  alt="no image"
                />
              </h1>
            </div>
          </div>
          <div className="w-1/2 h-full">
            <div className="w-full flex items-center p-2 justify-between border-[2px] border-gray rounded-t-sm">
              <p className="text-xl">Ngày/giờ ra:</p>
              <p className="text-xl font-bold">
                {exitResult?.created_at
                  ? dayjs(exitResult?.created_at).format("DD/MM/YYYY hh:mm:ss")
                  : ""}
              </p>
            </div>
            <div className="w-full shadow-xl py-2 px-2 border-[3px] border-gray rounded-b-sm">
              <p className="text-xl">Biển số xe ra</p>
              <img
                className={`w-full h-[100px] ${
                  exitResult?.crop_url ? "object-contain" : "object-cover"
                }`}
                src={exitResult?.crop_url ? exitResult.crop_url : no_image}
                alt="no image"
              />
            </div>
          </div>
        </div>
        <div className="w-1/2 h-full">
          <div className="w-full shadow-xl py-4 px-2 border-[2px] border-gray rounded-sm">
            <p className="text-xl">Giá vé</p>
            <h1 className="text-red-500 text-3xl font-bold text-center">
              {exitResult?.vehicle_type
                ? exitResult?.vehicle_type === "MOTORBIKE"
                  ? "4.0000D"
                  : "10.000D"
                : "Chưa xác định"}
            </h1>
          </div>
          <div className="w-full shadow-xl py-4 px-2 border-[2px] border-gray rounded-sm">
            <p className="text-xl">Loại xe</p>
            <h1 className="text-3xl font-bold text-center">
              {exitResult?.vehicle_type
                ? exitResult?.vehicle_type === "MOTORBIKE"
                  ? "Xe máy"
                  : "Xe ô tô"
                : "Chưa xác định"}
            </h1>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Detect;
