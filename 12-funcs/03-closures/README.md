# Closures as proxy

In this exercise, you will create a **on-off proxy** function.
The function should be called *proxy* and should satisfy the following criteria:
- it should receive function that has one string parameter and returns an int
- it should return a function that receives a string and returns an int and an error
- when the returned function is called, it should propagate the call to the received function and return its value and nil if the proxy is on, else it should return 0 and an error
- after every call, the state (on/off) is switched, the first time it's called it should be on


It is important that you do not externalize the state of the function, instead use closures to accomplish a similar effect.
Just for clarity, here is an example usage of the proxy function.
```python
  strlen := s -> len(s)
  decorated := proxy(strlen)
  decorated("alma") // returns (4,nil)
  decorated("xy") // returns (0,Error)
  decorated("xy") // returns (2,nil)
```

Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.