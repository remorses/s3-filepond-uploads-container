from aiohttp import web
from .logger import logger
import uuid
from aiobotocore.client import AioBaseClient
from .constants import S3_BASE_URL, S3_BUCKET, S3_DIR

routes = web.RouteTableDef()


@routes.post("/upload")
async def upload(request: web.Request):
    async for field in (await request.multipart()):
        if field.filename:
            data = await field.read(decode=True)
            key  = S3_DIR + str(uuid.uuid4())
            s3: AioBaseClient = request.app.create_s3()
            async with s3 as client:
                url = S3_BASE_URL + str(key)
                logger.debug(f"uploading file {url}")
                await client.put_object(
                    Bucket=S3_BUCKET, Key=key, Body=data, ACL="public-read"
                )
                return web.Response(text=url)
