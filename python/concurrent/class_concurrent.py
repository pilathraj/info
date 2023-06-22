
import concurrent.futures
import time


class Sum:
 def __init__(self):
  self.t1 = None
  self.t2 = None
  self.a = 0
  self.b = 10
 def get_a(self,a):
    time.sleep(3)
    self.a = a
    return a
 def get_b(self):
    time.sleep(3)
    self.b= 100
    return self.b

 def call(self):
   
   #concurrent
    with concurrent.futures.ThreadPoolExecutor(3) as my_executor:
        self.t1 = my_executor.submit(self.get_a, 10)
        self.t2 = my_executor.submit(self.get_b)
    # self.get_a(10)
    # self.get_b()
 
 def run(self):
      print(self.a + self.b)
      print(self.t1.result() +  self.t2.result())
     

start = time.perf_counter()
s = Sum()
s.call()
s.run()
finish = time.perf_counter()
print(f"total time {finish-start}") 
