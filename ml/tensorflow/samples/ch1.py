import tensorflow as tf
import numpy as np
print(tf.__version__) #2.12.0

#create a matrix
print("Matrix")
#m1
"""
[[1 1 1]
 [1 1 1]
 [1 1 1]]
"""
m1 = np.array([(1,1,1), (1,1,1), (1,1,1)], dtype="int32")
m2 = np.array([(2,2,2), (2,2,2), (2,2,2)], dtype="int32")

#m3
"""
[[3. 3. 3.]
 [3. 3. 3.]
 [3. 3. 3.]]
"""
m3 = np.array([(3,3,3), (3,3,3), (3,3,3)], dtype="float32")
m4 = np.array([(2,7,2),(1,4,2),(9,0,2)], dtype="float32")

print(m1)
print(m2)
print(m3)

#create a constant tensor form tensor like object
print("Constant tensor: ")
mc1 = tf.constant(m1)
mc2 = tf.constant(m2)

#mc3
"""
tf.Tensor(
[[3. 3. 3.]
 [3. 3. 3.]
 [3. 3. 3.]], shape=(3, 3), dtype=float32)
"""
mc3 = tf.constant(m3)

print(mc1)
print(mc2)
print(mc3)

mc_sum = tf.add(mc1, mc2)
print("Mc Sum: ", mc_sum) 
"""
Sum:  tf.Tensor(
[[3 3 3]
 [3 3 3]
 [3 3 3]], shape=(3, 3), dtype=int32)
"""

m_sum = tf.add(m1, m2)
print("M Sum: ", m_sum) 
"""
M Sum:  tf.Tensor(
[[3 3 3]
 [3 3 3]
 [3 3 3]], shape=(3, 3), dtype=int32)
"""

m_product = tf.matmul(m1, m3)
print("M product: ", m_product) 
"""
M product:  tf.Tensor(
[[9 9 9]
 [9 9 9]
 [9 9 9]], shape=(3, 3), dtype=int32)
"""

m4d = tf.linalg.det(m4)
print("matrix determinant M3: ", m4d) # matrix determinant M3:  tf.Tensor(55.999992, shape=(), dtype=float32)




