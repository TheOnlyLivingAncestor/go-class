# Sorting data

Write a function that receives a slice as its argument and returns a slice that contains the original values in ascending order (if the slice contains strings, sort from 'A' - 'Z').
The sorting should satisfy the following requirements:
- The function should work with slices of type *`uint8 and int32 }`* but is not required to work with other types
- You must not use type checking and the function must not accept a slice with type interface{} (hint: read the name of the parent directory :D)

An example for the usage:
```python
uints := {8, 128,2}
ints := {12, -3, 41, 27}
sortSlice(uints) // returns {2, 8, 128}
sortSlice(ints) // returns {-3, 12, 27, 41}
```


Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.