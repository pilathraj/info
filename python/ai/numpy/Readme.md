## Numpy - Numerical python
- Python library - working with arrays
- It has functions for  linear algebra, fourier transform, and matrices.

### Selling point
- 50x faster than Python list.
- NumPy aims to provide an array object that is 50x faster than a Python list.
- Array object in numpy is called 'ndarray'

### Why Numpy is faster?
- Numpy is stored at one continuous places in memory, unlike python list. So that process can access easily and we can manipulate data so easy.
- Most computation part in the numpy package has been written in C or C++, that will help us to improve the performance.
- 60% of code is written on Python.
- You can look at the source code at https://github.com/numpy/numpy
sample program:
```python
import numpy
arr = numpy.array([1, 2, 3, 4, 5])
print(arr)   # [1 2 3 4 5]
print(type(arr)) # <class 'numpy.ndarray'>
```
