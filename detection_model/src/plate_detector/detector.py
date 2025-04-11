from PIL import Image
import cv2
import torch
import src.plate_detector.function.utils_rotate as utils_rotate
import src.plate_detector.function.helper as helper
import numpy as np

class Model:
    def __init__(self):
        self.yolo_LP_detect = torch.hub.load('src/plate_detector/yolov5', 'custom', path='src/plate_detector/model/LP_detector.pt', force_reload=True, source='local')
        self.yolo_license_plate = torch.hub.load('src/plate_detector/yolov5', 'custom', path='src/plate_detector/model/LP_ocr.pt', force_reload=True, source='local')
        self.yolo_license_plate.conf = 0.60
        
    def __getPlates(self, img):
        
        plates = self.yolo_LP_detect(img, size=640)
        return plates
        
    def predict(self, img_file):
        img = cv2.imdecode(np.fromstring(img_file.read(), np.uint8), cv2.IMREAD_UNCHANGED)
        # If you use React-webcam for client
        img = cv2.flip(img, 1)
        plates = self.__getPlates(img)
        list_plates = plates.pandas().xyxy[0].values.tolist()
        list_read_plates = set()
        list_img_plate = []
        count = 0
        if len(list_plates) == 0:
            lp = helper.read_plate(self.yolo_license_plate, img)
            if lp != "unknown":
                list_read_plates.add(lp)
        else:
            for plate in list_plates:
                flag = 0
                x = int(plate[0]) # xmin
                y = int(plate[1]) # ymin
                w = int(plate[2] - plate[0]) # xmax - xmin
                h = int(plate[3] - plate[1]) # ymax - ymin
                crop_img = img[y:y+h, x:x+w]
                cv2.rectangle(img, (int(plate[0]),int(plate[1])), (int(plate[2]),int(plate[3])), color = (0,0,225), thickness = 2)
                # cv2.imwrite("crop.jpg", crop_img)
                # rc_image = cv2.imread("crop.jpg")
                list_img_plate.append(crop_img)
                lp = ""
                count+=1
                for cc in range(0,2):
                    for ct in range(0,2):
                        lp = helper.read_plate(self.yolo_license_plate, utils_rotate.deskew(crop_img, cc, ct))
                        if lp != "unknown":
                            list_read_plates.add(lp)
                            flag = 1
                            break
                    if flag == 1:
                        break
        return list_read_plates, list_img_plate
    
model = Model()