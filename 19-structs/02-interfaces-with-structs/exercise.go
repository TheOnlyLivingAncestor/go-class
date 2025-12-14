package structsinterfaces

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type Printable interface {
	Info() string
	PageNum() int
}

type Book struct {
	Author, Title string
	Pages         int
}

func (b Book) Info() string {
	return fmt.Sprintf("%v,%v", b.Author, b.Title)
}

func (b Book) PageNum() int {
	return b.Pages
}

func NewBook(Author, Title string, Pages int) Book {
	return Book{
		Author: Author,
		Title:  Title,
		Pages:  Pages,
	}
}

type Magazine struct {
	Title, Issue string
	Pages        int
}

func (m Magazine) Info() string {
	return fmt.Sprintf("%v,%v", m.Title, m.Issue)
}

func (m Magazine) PageNum() int {
	return m.Pages
}

func NewMagazine(Title, Issue string, Pages int) Magazine {
	return Magazine{
		Title: Title,
		Issue: Issue,
		Pages: Pages,
	}
}
