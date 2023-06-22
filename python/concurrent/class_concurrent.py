import concurrent.futures
import time


class Sum:
 def __init__():
  self.t1 = None
  self.t2 = None
  self.a = 0
 def get_a(self,a):
    time.sleep(3)
    self.a = a
    return a
 def get_b(self):
    time.sleep(3)
    self.b= 100
    return self.b

 def call(self):
    global t1, t2
    with concurrent.futures.ThreadPoolExecutor(3) as my_executor:
        t1 = my_executor.submit(get_a, 10)
        t2 = my_executor.submit(get_b)
 
 def run(self):
      print(t1.result() +  t2.result())
      print(self.a + self.b)

start = time.perf_counter()
s = sum()
s.call()
s.run()
finish = time.perf_counter()
print(f"total time {finish-start}") 
