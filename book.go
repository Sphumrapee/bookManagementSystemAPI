package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func searchAllBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func searchBookById(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func searchBooks(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	category := c.Query("category")

	fmt.Println("Searching for books with title:", title, ", author:", author, ", category:", category)
	var matchedBooks []Book

	for _, book := range books {
		if title != "" && !strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			continue
		}
		if author != "" && !strings.Contains(strings.ToLower(book.Author), strings.ToLower(author)) {
			continue
		}
		if category != "" && !strings.Contains(strings.ToLower(book.Category), strings.ToLower(category)) {
			continue
		}

		matchedBooks = append(matchedBooks, book)
	}

	if len(matchedBooks) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No books found matching the given filters",
		})
	}

	return c.JSON(matchedBooks)
}

func addBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	return c.JSON(book)
}

func editBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

func borrowBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			if book.IsBorrowed {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": "This book is already borrowed.",
				})
			}

			books[i].IsBorrowed = true
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("You have successfully borrowed '%s'.", book.Title),
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Book not found.",
	})
}

func returnBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			if !book.IsBorrowed {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": "This book is not borrowed.",
				})
			}

			books[i].IsBorrowed = false
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("You have successfully returned '%s'.", book.Title),
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Book not found.",
	})

}

func mostBorrowedBooks(c *fiber.Ctx) error {
	// Sort books by BorrowCount in descending order
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].BorrowCount > books[j].BorrowCount
	})

	// Return the most borrowed books (sorted)
	return c.JSON(books)
}
