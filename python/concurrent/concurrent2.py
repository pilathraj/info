import concurrent.futures
import time
# This not working at concurrently, seems i/o blocking. 
t1 = None
def get_a(a):
    time.sleep(1)
    return a
def get_b():
    time.sleep(1)
    print("get B")
    return 100

def call():
    global t1
    with concurrent.futures.ThreadPoolExecutor() as my_executor:
        t1 = my_executor.submit(get_b)
start = time.perf_counter()
call()
time.sleep(2)
print(get_a(20)+ t1.result())
finish = time.perf_counter()
print(f"total time {finish-start}") 



