import logging

import uvicorn
from fastapi import FastAPI

api = FastAPI()

logger = logging.getLogger(__name__)


@api.post("/")
async def index():
    pass


uvicorn.run(
    api,
    host="0.0.0.0",
    port=8080,
)
