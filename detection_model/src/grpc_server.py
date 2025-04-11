import sys
import os

sys.path.append(os.path.abspath(os.path.dirname(__file__)))
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import grpc
import cv2
import numpy as np
import time
import src.proto.gen.detect_pb2 as pb2
import src.proto.gen.detect_pb2_grpc as pb2_grpc
from concurrent import futures
from src.plate_detector.detector import model
from minio import Minio
from minio.error import S3Error
import io

# Cấu hình MinIO
MINIO_ENDPOINT = os.getenv("MINIO_ENDPOINT", "parking.minio:9000")
MINIO_ACCESS_KEY = os.getenv("MINIO_ACCESS_KEY", "3SYhDzVQBrLI9SzRB1zR")
MINIO_SECRET_KEY = os.getenv("MINIO_SECRET_KEY", "8LJFPwOg4jscApFpAwawnbTKHNcyTd6y60mOzZbs")
BUCKET_NAME = "goparking"
FOLDER_NAME = "crop_img"

# Khởi tạo MinIO client
minio_client = Minio(
    MINIO_ENDPOINT,
    access_key=MINIO_ACCESS_KEY,
    secret_key=MINIO_SECRET_KEY,
    secure=False  # For HTTP or Docker Local
    # secure=True  # For HTTPS
)

#For HTTPS
# if not minio_client.bucket_exists(BUCKET_NAME):
#     minio_client.make_bucket(BUCKET_NAME)

def upload_to_minio(file_path, file_name):
    try:
        object_name = f"{FOLDER_NAME}/{file_name}"
        minio_client.fput_object(BUCKET_NAME, object_name, file_path)
        # file_url = f"https://{MINIO_ENDPOINT}/{BUCKET_NAME}/{object_name}" For HTTPS
        file_url = f"http://localhost:9000/{BUCKET_NAME}/{object_name}"
        return file_url
    except S3Error as e:
        print(f"MinIO Error: {e}")
        return None

class PlateDetectionServicer(pb2_grpc.PlateDetectionServicer):
    def Detect(self, request, context):
        # Convert `bytes` to file-like object
        img_file = io.BytesIO(request.image)

        plate_list, crop_img_list = model.predict(img_file) 

        if len(plate_list) != 1:
            return pb2.PlateResponse(
                license_plate_detect=[],
                crop_img_url=""
            )

        crop_url = ''
        if len(crop_img_list) == 1:
            # Đặt tên file duy nhất theo timestamp
            crop_file_name = f"crop_{int(time.time())}.jpg"
            crop_file_path = crop_file_name

            cv2.imwrite(crop_file_path, crop_img_list[0])

            # Upload to MinIO
            crop_url = upload_to_minio(crop_file_path, crop_file_name)

            os.remove(crop_file_path)

        return pb2.PlateResponse(
            license_plate_detect=list(plate_list),
            crop_img_url=crop_url
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    pb2_grpc.add_PlateDetectionServicer_to_server(PlateDetectionServicer(), server)
    server.add_insecure_port('0.0.0.0:50051')
    print("gRPC server is running on port 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
