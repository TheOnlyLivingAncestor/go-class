# Interfaces with Structs

In this exercise, you'll build a game store along with a set of functions to query the store.

1. Define a `Printable` interface with the following functions:
   - `Info()` provides all available textual information (i.e., members except the page number) comma-separated and returns `string`
   - `PageNum()` shows number of pages and returns `int`

2. Declare the following structs with appropriate elements:
   - `Book`
   - `Magazine`

3. Write a constructor functions:
   - `NewBook(Author,Title string, Pages int) Book`
   - `NewMagazine(Title,Issue string, Pages int) Magazine`

3. Implement interface `Printable` for `Book` and `Magazine`

Insert your code into the file `exercise.go`.

Hint: read about how to use [interfaces](https://go.dev/tour/methods/9).
