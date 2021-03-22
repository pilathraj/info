# Strings
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

