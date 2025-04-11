# import cv2
# import os
# import time
# from flask import Flask, request, jsonify
# from src.plate_detector.detector import model  
# # import cloudinary.uploader // if you want to use cloudinary
# from flask_cors import CORS
# from minio import Minio
# from minio.error import S3Error

# MINIO_ENDPOINT = os.getenv("MINIO_ENDPOINT", "139.59.250.218:9000")
# MINIO_ACCESS_KEY = os.getenv("MINIO_ACCESS_KEY", "1d1KqQif0abfpxQaxyy0")
# MINIO_SECRET_KEY = os.getenv("MINIO_SECRET_KEY", "i9n4WD3PFSzc6XFZkow69UWV5dGx6bZbmMIjaNlU")
# BUCKET_NAME = "goparking"
# FOLDER_NAME="crop_img"

# # Khởi tạo MinIO client
# minio_client = Minio(
#     MINIO_ENDPOINT,
#     access_key=MINIO_ACCESS_KEY,
#     secret_key=MINIO_SECRET_KEY,
#     # secure=False  # For HTTP
#     secure=True  # For HTTPS
# )

# print(MINIO_ENDPOINT,MINIO_ACCESS_KEY,MINIO_SECRET_KEY,BUCKET_NAME,)

# if not minio_client.bucket_exists(BUCKET_NAME):
#     minio_client.make_bucket(BUCKET_NAME)

# def upload_to_minio(file_path, file_name):
#     try:
#         object_name = f"{FOLDER_NAME}/{file_name}"
        
#         # Upload file
#         minio_client.fput_object(BUCKET_NAME, object_name, file_path)

#         # Tạo URL truy cập file
#         file_url = f"https://{MINIO_ENDPOINT}/{BUCKET_NAME}/{object_name}"
#         return file_url
#     except S3Error as e:
#         print(f"MinIO Error: {e}")
#         return None

# def create_app():
#     app = Flask(__name__)
#     CORS(app)

#     @app.route('/detect', methods=['POST'])
#     def classify():
#         license_plate = request.files.get('image')

#         if not license_plate:
#             return {"error": "No image provided"}, 400

#         # Gọi model để dự đoán biển số
#         plate_list, crop_img_list = model.predict(license_plate) 
#         print(len(plate_list))

#         if len(plate_list) != 1:
#             return jsonify({
#                 "message": "There are not one plate" if len(plate_list) == 0 else "There are more one plate",
#             })

#         # Nếu có một biển số hợp lệ, xử lý ảnh cắt được
#         crop_url = ''
#         if len(crop_img_list) == 1:
#             # Đặt tên file duy nhất theo timestamp
#             crop_file_name = f"crop_{int(time.time())}.jpg"
#             crop_file_path = crop_file_name

#             # Lưu ảnh tạm trên server
#             cv2.imwrite(crop_file_path, crop_img_list[0])

#             # Upload ảnh lên MinIO
#             crop_url = upload_to_minio(crop_file_path, crop_file_name)

#             # Xóa file tạm sau khi upload
#             os.remove(crop_file_path)

#         return jsonify({
#             "license_plate_detect": list(plate_list),
#             "crop_img_url": crop_url,
#         })

#     return app


# ## if you want to use cloudinary
#     # cv2.imwrite("crop.jpg", crop_img_list[0])
#     # crop_response=cloudinary.uploader.upload(
#     # 'crop.jpg',
#     # folder="parking",
#     # unique_filename = True, 
#     # overwrite=True,
#     # eager=[{"width": 200, "crop": "fill"}])
#     # crop_url = crop_response["url"]
#     # os.remove('crop.jpg')