import concurrent.futures
import time

t1, t2 = None, None
def get_a(a):
    time.sleep(3)
    return a
def get_b():
    time.sleep(3)
    return 100

def call():
    global t1, t2
    with concurrent.futures.ThreadPoolExecutor(3) as my_executor:
        t1 = my_executor.submit(get_a, 10)
        t2 = my_executor.submit(get_b)
         

start = time.perf_counter()
call()
print(t1.result() +  t2.result())
finish = time.perf_counter()
print(f"total time {finish-start}") 



