package main

import (
	"github.com/gofiber/fiber/v2"
)

// Book struct to hold book data
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	IsBorrowed  bool   `json:"is_borrowed"`
	BorrowCount int    `json:"borrow_count"`
}

var books = []Book{
	{ID: 1, Title: "Go Programming", Author: "John Doe", Category: "Programming", IsBorrowed: false, BorrowCount: 2},
	{ID: 2, Title: "Go Programming", Author: "Jane Doe", Category: "Programming", IsBorrowed: true, BorrowCount: 3},
	{ID: 3, Title: "Learning Fiber", Author: "Alice Smith", Category: "Web Development", IsBorrowed: false, BorrowCount: 5},
	{ID: 4, Title: "Advanced Go", Author: "John Doe", Category: "Programming", IsBorrowed: false, BorrowCount: 9},
}

func main() {
	app := fiber.New()

	app.Get("/books", searchAllBooks)
	app.Get("/books/search/id/:id", searchBookById)
	app.Get("/books/search", searchBooks)
	app.Post("/books/create", addBook)
	app.Put("/books/Update/:id", editBook)
	app.Delete("/books/deleteBook/:id", deleteBook)
	app.Post("/books/borrow/:id", borrowBook)
	app.Post("/books/return/:id", returnBook)
	app.Get("/books/most-borrowed", mostBorrowedBooks)

	app.Listen(":8080")
}
