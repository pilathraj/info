# 1. Basics
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
# 2. Strings
## 1. Index Slice
```python
n = "0123456789"
print(n) #0123456789
print(n[1]) # 1  -> index start from Zero
print(n[1:3]) # 12  -> Start and end-1 index, Print from 1 to 2 index 
print(n[6:]) # 6789
print(n[:5]) # 01234
print(n[-5:-2]) # 567
```
## String function
```python
s = "Pilathraj"
print(len(s)) #9
#  Substring check
print("raj" in s) # True
# print array 
for c in "cse":
  print(c)  
  """ 
  output:
  c
  s
  e
  """
print(s.upper()) # PILATHRAJ
print(" Hello Pilathraj ".strip()) # trim the whitespace Hello Pilathraj
print(s.count("a")) # 2 <- how many times substring print in the string s.
print(s.find("raj")) # 6 <- retrun starting position
print("Hello World!!!".replace("World", "Pilath")) # Hello Pilathaj!!!
print("Hello World!!!".split(" ")) # ['Hello', 'World!!!']
print("My name is {}, I'm {} years old".format("Mark", 45)) # My name is Mark, I'm 45 years old
quantity = 3
itemno = "mobile"
price = 49.95
myorder = "I want to pay {2} dollars for {0} pieces of item {1}."
print(myorder.format(quantity, itemno, price)) #I want to pay 49.95 dollars for 3 pieces of item mobile.
txt = "apple, banana, cherry"
print(txt.split(", ", 1)) # ['apple', 'banana, cherry']
print(txt.rsplit(", ", 1)) # ['apple, banana', 'cherry'] <- right to left
print("50".zfill(10)) # 0000000050
```
# 3. Collections
- **List** is a collection which is ordered and changeable. Allows duplicate members.
- **Tuple** is a collection which is ordered and unchangeable. Allows duplicate members.
- **Set** is a collection which is unordered and unindexed. No duplicate members.
- **Dictionary** is a collection which is unordered and changeable. No duplicate members.
## List
```python 
list1 = ['mark', 26, 'First street', True]
list2 = [1,2,3]
list3 = []
list4 = list2 # copy the reference
list5 = list2.copy() # copy the values
# Change the item

list2[2] = 4
print(list2) # [1, 2, 4]
print("list4:", list4) # [1,2,4]
print("list5: ", list5) # [1,2,3]
list2[1:] = [2,3] 
print("list2:", list2) # [1, 2, 3]

# Join the list
print(list1 + list2) # ['mark', 26, 'First street', True, 1, 2, 3]
for i in list2:
	list3.append(i)
print("list3", list3) # list3: [1, 2, 3]
list1.extend(list2) # Copied items from list2 to list1 
print(list1) # ['mark', 26, 'First street', True, 1, 2, 3]

# Remove item
list1.remove(True)
print(list1) # ['mark', 26, 'First street', 1, 2, 3]
list1.pop(-2) 
print(list1) # ['mark', 26, 'First street', 1, 3]
list3.clear() # remove all the elements
print("list3", list3) # []
```
- **List Comprehension**
  - List comprehension offers a shorter syntax when you want to create a new list based on the values of an existing list.
```python
fruits = ["apple", "banana", "cherry", "kiwi", "mango"]
newlist = [x for x in fruits if "a" in x]
print(newlist) # ['apple', 'banana', 'mango']
newlist = [x for x in fruits]
print(newlist) # ['apple', 'banana', 'cherry', 'kiwi', 'mango']
newlist = [x if x != "banana" else "orange" for x in fruits]
print(newlist) #['apple', 'orange', 'cherry', 'kiwi', 'mango']
```
- **Sort list**
```python
#Case sensitive sorting can give an unexpected result:
thislist = ["banana", "Orange", "Kiwi", "cherry"]
thislist.sort()
print(thislist) # ['Kiwi', 'Orange', 'banana', 'cherry']

#Case sensitive sorting descending order
thislist = ["banana", "Orange", "Kiwi", "cherry"]
thislist.sort(reverse = True)
print(thislist)  # ['cherry', 'banana', 'Orange', 'Kiwi']

thislist = ["banana", "Orange", "Kiwi", "cherry"]
thislist.sort(key = str.lower) # Case insensitive ascending order
print(thislist) # ['banana', 'cherry', 'Kiwi', 'Orange']

thislist = ["banana", "Orange", "Kiwi", "cherry"]
thislist.sort(key = str.lower, reverse = True) # Case insensitive descending order
print(thislist) # ['Orange', 'Kiwi', 'cherry', 'banana']

thislist = ["banana", "Orange", "Kiwi", "cherry"]
thislist.reverse() # print same list in the reverse order
print(thislist) # ['cherry', 'Kiwi', 'Orange', 'banana']
```
## Tuples
- Tuples are read-only ordered items.
- Tuples are written with round brackets.

```python 
   t1 = ("abc", 34, True, 40, "male")
   print(t1) # ('abc', 34, True, 40, 'male')
   thistuple = ("apple",)
print(type(thistuple)) # class tuple

#NOT a tuple
thistuple = ("apple")
print(type(thistuple)) # class string
 ```  
