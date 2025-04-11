import sys
import os
# import cloudinary
# import cloudinary.uploader
# from dotenv import load_dotenv 

sys.path.append(os.path.abspath(os.path.dirname(__file__)))
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))


# load_dotenv()
# cloudinary.config(
#     cloud_name=os.getenv("CLOUDINARY_CLOUD_NAME"),
#     api_key=os.getenv("CLOUDINARY_API_KEY"),
#     api_secret=os.getenv("CLOUDINARY_API_SECRET")
# )

from src import create_app

application = create_app()

if __name__ == "__main__":
    application.run(host="0.0.0.0", port=5000)

