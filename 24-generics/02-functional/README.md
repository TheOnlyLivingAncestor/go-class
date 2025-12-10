# Generic functions

In this exercise you  will have to implement a generic *filter* function.
The filter function receives as arguments a slice of type E and a function of type E -> bool,
and returns a slice of type E.

Filter is a function that recieves a list and a predicate function and returns  a list containing the elements from the input list where the predicate returns true.
For example:
```python
    is_even := a -> a % 2 is 0
    values := {1,2,3,4,5}
    filtered := filter(values, is_even) // produces {2,4}
```

Place your code into the file exercise.go near the placeholder // INSERT YOUR CODE HERE.