import os


AWS_ACCESS_KEY_ID = os.getenv("ACCESS_KEY_ID")
AWS_SECRET_ACCESS_KEY = os.getenv("SECRET_ACCESS_KEY")
# S3_ENDPOINT = "https://fra1.digitaloceanspaces.com"
S3_DIR = os.getenv("DIRECTORY", "")
# S3_BASE_URL = "https://instagrammedias.fra1.cdn.digitaloceanspaces.com/"
S3_BUCKET = os.getenv("BUCKET")
S3_REGION_NAME = os.getenv("REGION")
S3_BASE_URL = f"https://{S3_BUCKET}.s3-{S3_REGION_NAME}.amazonaws.com/"


assert all([AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, S3_BUCKET, S3_REGION_NAME])

TESTING = bool(os.getenv("TESTING"))
JWT_SECRET = os.getenv("JWT_SECRET", "secret")
JWT_ALGORITHM = "HS256"
CLIENT_ID = os.getenv("GOOGLE_CLIENT_ID")
CLIENT_SECRET = os.getenv("GOOGLE_CLIENT_SECRET")
