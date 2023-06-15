import concurrent.futures
import time



def get_a(a):
    time.sleep(3)
    return a
def get_b():
    time.sleep(3)
    return 100

def sum():
    sum = get_a(10) + get_b()
    return sum


start = time.perf_counter()
print(sum())
finish = time.perf_counter()
print(f"total time {finish-start}") 


