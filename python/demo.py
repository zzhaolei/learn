import asyncio
import tracemalloc


async def sub_demo() -> None:
    for i in range(10):
        await asyncio.sleep(0.1)


async def benchmark(num: int) -> float:
    tracemalloc.start()
    tasks = [asyncio.create_task(sub_demo()) for _ in range(num)]
    await asyncio.sleep(0)
    snapshot = tracemalloc.take_snapshot()
    await asyncio.wait(tasks)  # 等待执行完成，防止影响到后面的其他测试
    total_bytes = sum(stat.size for stat in snapshot.statistics('lineno'))
    return total_bytes / 1024.0


async def main() -> None:
    for n in [2000, 5000, 10000, 50000, 100000]:
        total_kb = await benchmark(n)
        print(f'> coroutines={n:5} used {total_kb:.3f} K/b  per:{total_kb / n} K/b')


asyncio.run(main())