package basics

import (
	"log"
)

type WithId interface {
	GetId() int64
}

type GenericCrudDb[T WithId] interface {
	GetById(id int64) (*T, bool)
	Add(book T) (T, error)
	DeleteById(id int64) (*T, error)
}

func InitBookCrudDb() GenericCrudDb[Book] {
	return InitBookDb()
}

func Generics() {
	bookCrudDb := InitBookCrudDb()
	addedBook, err := bookCrudDb.Add(NewBook("Some Book", "Some Author", 9999999999, 1337))
	if err != nil {
		log.Fatal(err)
	}
	addedBook.PrintBook()
}
