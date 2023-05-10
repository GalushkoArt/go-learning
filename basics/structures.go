package basics

import "fmt"

type ISBN int64

type Book struct {
	id     int64
	author string
	title  string
	pages  int
	isbn   ISBN
	used   int
}

func (b Book) GetId() int64 {
	return b.id
}

func NewBook(title string, author string, isbn int64, pages int) Book {
	return Book{
		title:  title,
		author: author,
		isbn:   ISBN(isbn),
		pages:  pages,
		used:   0,
	}
}

func (b *Book) PrintBook() {
	b.used++
	fmt.Printf("%+v\n", b)
}

func Structures() {
	request := struct {
		method string
		url    string
		body   string
	}{"GET", "/index.html", ""}
	fmt.Printf("%+v\n", request)
	cleanCode := NewBook("Clean Code", "Uncle Bob", 94357234982374, 462)
	fmt.Println(cleanCode.title, cleanCode.author, cleanCode.pages, cleanCode.isbn)
	cleanCode.PrintBook()
}
