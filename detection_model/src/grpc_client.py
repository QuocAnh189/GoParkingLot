import sys
import os

sys.path.append(os.path.abspath(os.path.dirname(__file__)))
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import grpc
import cv2
import numpy as np
import src.proto.gen.detect_pb2 as pb2
import src.proto.gen.detect_pb2_grpc as pb2_grpc

def run_client(image_path):
    img = cv2.imread(image_path)

    # Chuyển đổi ảnh thành bytes (gRPC không hỗ trợ numpy.ndarray)
    _, img_encoded = cv2.imencode('.jpg', img)
    img_bytes = img_encoded.tobytes()

    # Tạo gRPC channel
    channel = grpc.insecure_channel('localhost:50051')
    stub = pb2_grpc.PlateDetectionStub(channel)

    # Gửi request với ảnh dưới dạng bytes
    request = pb2.PlateRequest(image=img_bytes)
    response = stub.Detect(request)

    print("Response from server:", response.license_plate_detect, response.crop_img_url)

if __name__ == "__main__":
    image_path = "src/plate_license.jpg"
    run_client(image_path)