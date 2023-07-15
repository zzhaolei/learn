import time
from queue import Queue
from threading import Thread
from typing import Any

queue = Queue()


class Pool:
    def __init__(self, max_workers: int, max_queues: int):
        self.queue = Queue(maxsize=max_queues)
        self._start_threads(max_workers)

    def _start_threads(self, max_workers: int):
        for i in range(max_workers):
            t = Thread(target=self.do_job)
            t.daemon = True
            t.start()

    def submit(self, job: Any, block=True, timeout=None):
        return self.queue.put(job, block=block, timeout=timeout)

    def do_job(self):
        while True:
            job = self.queue.get()
            print(f"this is job {job}")
            self.queue.task_done()

    def wait(self):
        self.queue.join()


pool = Pool(4, 4)
pool.submit("this is 1")
pool.submit("this is 2")
time.sleep(2)
print("sleep done")
pool.wait()
