package structembedding

import "encoding/json"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// Author represents information about the book's author
type Author struct {
	Name, Address string
}

// Book represents information about a book
type Book struct {
	Author      Author
	Title, ISBN string
	Pages       int
}

// Article represents information about a article
type Article struct {
	Author         Author
	Title, Journal string
	Year           int
}

// ParseBook parses the given JSON data into a Book struct
func ParseBook(jsonData []byte) (Book, error) {
	var b Book
	err := json.Unmarshal(jsonData, &b)
	if err != nil {
		return Book{}, err
	}
	return b, nil
}

// ParseArticle parses the given JSON data into a Article struct
func ParseArticle(jsonData []byte) (Article, error) {
	var a Article
	err := json.Unmarshal(jsonData, &a)
	if err != nil {
		return Article{}, err
	}
	return a, nil
}
