## Numpy - Numerical python
- Python library - working with arrays
- It has functions for  linear algebra, Fourier transform, and matrices.

### Selling point
- 50x faster than Python list.
- NumPy aims to provide an array object that is 50x faster than a Python list.
- Array object in numpy is called 'ndarray'

### Why Numpy is faster?
- Numpy is stored at one continuous place in memory, unlike Python list. So that process can access easily and we can manipulate data so easy.
- Most computation part in the numpy package has been written in C or C++, which will help us to improve the performance.
- 60% of code is written in Python.
- You can look at the source code at https://github.com/numpy/numpy
sample program:
```python
import numpy as np
arr = np.array([1, 2, 3, 4, 5])
print(arr)   # [1 2 3 4 5]
print(type(arr)) # <class 'numpy.ndarray'>
print(np.__version__)
```

### create multi dimension array
Sample program:
```python
import numpy as np

a = np.array(42)
b = np.array([1, 2, 3, 4, 5])
c = np.array([[1, 2, 3], [4, 5, 6]])
d = np.array([[[1, 2, 3], [4, 5, 6]], [[1, 2, 3], [4, 5, 6]]])

print(a.ndim)
print(b.ndim)
print(c.ndim)
print(d.ndim)
```
