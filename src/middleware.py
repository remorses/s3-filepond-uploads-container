import jwt
import aiohttp.web
from .logger import logger
from .constants import JWT_SECRET, JWT_ALGORITHM, TESTING


async def jwt_middleware(app, handler):
    async def middleware(request: aiohttp.web.Request):
        if request.method in ["OPTIONS"]:
            return await handler(request)
        request.jwt = {}
        jwt_token = request.headers.get("Authorization", "").replace("Bearer ", "")
        if jwt_token:
            try:
                payload = jwt.decode(
                    jwt_token, key=JWT_SECRET, verify=True, algorithms=[JWT_ALGORITHM]
                )
            except (jwt.DecodeError, jwt.ExpiredSignatureError) as e:
                # return to_response(*fail('Token is invalid'))
                logger.error(e)
                return await handler(request)
            else:
                request.jwt = payload
        else:
            logger.debug("no Authorization")
        # if TESTING and not request.jwt:
        #     logger.warn("using TESTING fake user_id")
        # request.jwt = {"user_id": "000000000000000000000000"}
        return await handler(request)

    return middleware
