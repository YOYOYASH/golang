package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(book Book) (Book,error) {
	if book.Copies == 0 {
		return Book{}, errors.New("no copies of this book is curently available")
	}
	book.Copies--
	return book,nil
}
