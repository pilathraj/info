## Basics
```python
#Variable Declaration
a = "test"
x, y, z, n = "Orange", "Banana", "Cherry", 5
s = t = v = "Orange"
i,f,cx = 10, 10.10, 10+10j
# Single line command

""""
Multi
line 
cmd"""

# array === list
a = [1,2,3] # override
# Unpack a list:
b,c,d = a

print(x)
print(y, end=" ")
print(z)
print(n)
print(a)
print(b,c,d)
print(s+t+v) #Concatenation Operator
print(i,f,cx, end= ": ")
print(type(i),type(f),type(cx))
```
## Output
```
Orange
Banana Cherry
5
[1, 2, 3]
1 2 3
OrangeOrangeOrange
10 10.1 (10+10j): <class 'int'> <class 'float'> <class 'complex'>
```
### Variable scoping
```python
x = "awesome"

def myfunc():
  global x, z
  x = "fantastic"
  y = "testing"
  z = "global variable create inside a function"

myfunc()

print("x: " + x)
#print("Access local variable " + y) #NameError: name 'y' is not defined
print("z: " + z)
#print("Access undefined variables ", xyz) #NameError: name 'xyz' is not defined

```
### output
```
x: fantastic
z: global variable created inside a function
```
### Data types
|Description| type| example|
|-----------|-----|--------|
|Text Type:|	str| s = "Pilath"  (or) S= 'raj' or MS="""Multi line string"""|
|Numeric Types:|	int, float, complex | i,f,c = 10, 10.10, 10+10j|
|Sequence Types:|	list, tuple, range| l, t, r =["apple", "banana", "cherry"], ("apple", "banana", "cherry"), range(1,5)|
|Mapping Type:|	dict| x = {"name" : "John", "age" : 36}|
|Set Types:|	set, frozenset| s, f =  {"apple", "banana", "cherry"}, frozenset({"apple", "banana", "cherry"})|
|Boolean Type:|	bool| b =True|
|Binary Types:|	bytes, bytearray, memoryview| b, ba, mv = b"Hello", bytearray(5), memoryview(bytes(5))|

