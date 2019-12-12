import asyncio
import os
import traceback
import urllib.parse

import aiobotocore
import aiohttp_cors
from aiobotocore.client import AioBaseClient
from aiohttp import web

from .constants import AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, S3_REGION_NAME # S3_ENDPOINT

from .logger import logger
from .middleware import jwt_middleware
from .routes import routes


def build():
    app = web.Application(middlewares=[jwt_middleware], )
    session = aiobotocore.get_session()
    app.create_s3 = lambda: session.create_client(
        "s3",
        region_name=S3_REGION_NAME,
        # endpoint_url=S3_ENDPOINT, # TODO support other storage via endpoint url
        aws_secret_access_key=AWS_SECRET_ACCESS_KEY,
        aws_access_key_id=AWS_ACCESS_KEY_ID,
    )
    app.add_routes(routes)
    cors = aiohttp_cors.setup(
        app,
        defaults={
            "*": aiohttp_cors.ResourceOptions(
                allow_credentials=True,
                expose_headers="*",
                allow_headers="*",
                allow_methods="*",
            )
        },
    )
    for route in list(app.router.routes()):
        cors.add(route)

    return app


if __name__ == "__main__":
    web.run_app(build(), port=80, shutdown_timeout=6.)
