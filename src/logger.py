import coloredlogs
import os
import logging

logger = logging.getLogger(__name__)
LOGGING_FORMAT = os.getenv('LOGGING_FORMAT') or "[%(asctime)s] [%(levelname)s] %(message)s"

coloredlogs.DEFAULT_FIELD_STYLES = {
    "asctime": {"color": "white"},
    "hostname": {"color": "white"},
    "levelname": {"color": "white", "bold": True},
    "name": {"color": "white"},
    "programname": {"color": "white"},
}
coloredlogs.install(
    fmt=LOGGING_FORMAT,
    level="DEBUG",
    logger=logger,
)
