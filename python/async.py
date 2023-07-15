import asyncio
from time import process_time as time

import uvloop


async def time_tasks(count: int =100) -> float:

    async def nop_task() -> None:
        pass

    start = time()
    tasks = [asyncio.create_task(nop_task()) for _ in range(count)]
    await asyncio.gather(*tasks)   # <--移动在elapsed之前
    elapsed = time() - start
    return elapsed


async def main() -> None:
    for count in range(100_000, 1000_000 + 1, 100_000):
        create_time = await time_tasks(count)
        create_per_second = 1 / (create_time / count)
        print(f"{count:,} tasks \t {create_per_second:0,.0f} tasks per/s")

asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
asyncio.run(main())