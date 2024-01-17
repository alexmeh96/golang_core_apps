package main

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var books = []*Book{
	{
		ID:               "1",
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

func listBooks() []*Book {
	return books
}

func getBook(id string) *Book {
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func storeBook(book Book) {
	books = append(books, &book)
}

func deleteBook(id string) *Book {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], (books)[i+1:]...)
			return &Book{}
		}
	}
	return nil
}

func updateBook(id string, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = &bookUpdate
			return book
		}
	}
	return nil
}
