# Count lowercase characters

Write a function that reads the below text and returns the number of "lowercase characters".

> "Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts. Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean. A small river named Duden flows by their place and supplies it with the necessary regelialia. It is a paradisematic country, in which roasted parts of sentences fly into your mouth."

The signature of the function should be `func counter(io.Reader) int`. The argument is an [`io.Reader`](https://pkg.go.dev/io#Reader), a basic I/O interface that supports a single method to read the above input text:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

You can solve the exercise by using the barebones `io.Reader` interface, but this will be a lot of work. To simplify dealing with buffered reads, Go provides lots of useful convenience packages in the standard library, like [`bufio.ReadString`](https://pkg.go.dev/bufio#Reader.ReadString) to read a buffer into a string until a delimiter (use [`bufio.NewReader`](https://pkg.go.dev/bufio#NewReader) to obtain a `bufio.Reader` from an `io.Reader`), [`strings.Fields`](https://pkg.go.dev/strings#Fields) to split a string into a slice of strings on whitespace characters, and [`unicode.IsLower`](https://pkg.go.dev/unicode#IsLower) and [`unicode.IsUpper`](https://pkg.go.dev/unicode#IsUpper) to check whether a Unicode character is upper-case or lowercase, respectively. Use your best judgement to assemble these helpers into your own solution.

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.
