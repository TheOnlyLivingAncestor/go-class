# Return the directory attribute from a path

Write a function to split a file system path into a directory and filename component, return the
directory part and discard the other component. Use the `/` symbol as a directory
separator. If the path consists of a single filename with no directory, then return an empty directory.


For instance, the directory part of the path `/usr/bin/go` is
`/usr/bin`, and the directory part of `vmlinuz` is
``

Implement your code in a function called `splitPath` that has the below signature:

``` go
// splitPath returns the directory component of a file path.
func splitPath(fullPath string) string {
    ...
}
```

Insert your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.

HINT: use the [`path.Split()`](https://pkg.go.dev/path#Split) function.
