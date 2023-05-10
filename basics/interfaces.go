package basics

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BookCrudDb interface {
	GetById(id int64) (*Book, bool)
	Add(book Book) (Book, error)
	DeleteById(id int64) (*Book, error)
}

type BookSearch interface {
	GetByTitle(title string) []Book
	GetByAuthor(author string) []Book
}

type BookDb interface {
	BookCrudDb
	BookSearch
}

type MapMockBookCrudDb struct {
	sequence int64
	storage  map[int64]*Book
}

func InitBookDb() BookDb {
	return MapMockBookCrudDb{storage: make(map[int64]*Book)}
}

func (mock MapMockBookCrudDb) Add(book Book) (Book, error) {
	mock.sequence++
	book.id = mock.sequence
	mock.storage[book.id] = &book
	return book, nil
}

func (mock MapMockBookCrudDb) DeleteById(id int64) (*Book, error) {
	book, ok := mock.GetById(id)
	if !ok {
		return nil, errors.New("book with " + strconv.FormatInt(id, 10) + " id was not found")
	}
	delete(mock.storage, id)
	return book, nil
}

func (mock MapMockBookCrudDb) GetById(id int64) (*Book, bool) {
	book, ok := mock.storage[id]
	return book, ok
}

func (mock MapMockBookCrudDb) GetByTitle(title string) []Book {
	search := strings.ToLower(title)
	results := make([]Book, 0, 5)
	for _, book := range mock.storage {
		if strings.Contains(strings.ToLower(book.title), search) {
			results = append(results, *book)
		}
	}
	return results
}

func (mock MapMockBookCrudDb) GetByAuthor(author string) []Book {
	search := strings.ToLower(author)
	results := make([]Book, 0, 5)
	for _, book := range mock.storage {
		if strings.Contains(strings.ToLower(book.author), search) {
			results = append(results, *book)
		}
	}
	return results
}

func Interfaces() {
	mockDb := InitBookDb()
	book, err := mockDb.Add(NewBook("DDD", "Eric Evans", 94324834578, 403))
	if err != nil {
		log.Fatal(err)
	}
	found, ok := mockDb.GetById(book.id)
	if ok {
		found.PrintBook()
	} else {
		log.Fatal("Failed to save the Book!")
	}
	foundByAuthor := mockDb.GetByAuthor("Evans")
	fmt.Printf("%+v\n", foundByAuthor)
	foundByTitle := mockDb.GetByTitle("ddd")
	fmt.Printf("%+v\n", foundByTitle)
	printInterface(Summer)
	printInterface(byte(255))
	printInterface("Hello! It's me!")
}

func printInterface(i interface{}) {
	switch i.(type) {
	case Season:
		fmt.Println("Wow! It's season", i)
	case byte:
		fmt.Println("Why did you pass byte?")
		fmt.Println(strconv.FormatInt(int64(i.(byte)), 16))
	default:
		fmt.Println(i)
	}
}
