# Search

Write a function that reads the below text and returns `true` or `false` if "TIME" appears in the text *more than once*, without respect to letter case (i.e., the each of words `Hello`, `heLLo`, and `HELlo` match for the search strings `HELLO` and `hello`).

> "The great secret known to Apollonius of Tyana, Paul of Tarsus, Simon Magus, Asklepios, Paracelsus, Boehme and Bruno is that: we are moving backward in time. The universe in fact is contracting into a unitary entity which is completing itself. Decay and disorder are seen by us in reverse, as increasing. These healers learned to move forward in time, which is retrograde to us."

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

HINT:

- Make sure to remove punctuation from the input (like `.`, `;`, or `?`) and trim the text to alphanumeric characters. For instance, you can use [`regexp.ReplaceAllString`](https://pkg.go.dev/regexp#Regexp.ReplaceAllString) with the regular expression `[^a-zA-Z0-9 ]+` for this purpose.
- You can use a `map[string]int` to remember the number of times a particular word is found while iterating through the text word by word.
