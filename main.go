package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

var books = []Book{
	{ID: 1, Title: "The Alchemist", Author: "Paulo Coelho", Publisher: "HarperCollins"},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Publisher: "J. B. Lippincott & Co."},
	{ID: 3, Title: "1984", Author: "George Orwell", Publisher: "Secker & Warburg"},
}

func main() {
	app := fiber.New()

	// Get all books
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	// Get a book by ID
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, book := range books {
			if book.ID == atoi(id) {
				return c.JSON(book)
			}
		}
		return c.Status(404).SendString("Book not found")
	})

	// Add a new book
	app.Post("/books", func(c *fiber.Ctx) error {
		var book Book
		if err := c.BodyParser(&book); err != nil {
			return err
		}
		book.ID = len(books) + 1
		books = append(books, book)
		return c.JSON(book)
	})

	// Update a book by ID
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var newBook Book
		if err := c.BodyParser(&newBook); err != nil {
			return err
		}
		for i, book := range books {
			if book.ID == atoi(id) {
				newBook.ID = book.ID
				books[i] = newBook
				return c.JSON(newBook)
			}
		}
		return c.Status(404).SendString("Book not found")
	})

	// Delete a book by ID
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, book := range books {
			if book.ID == atoi(id) {
				books = append(books[:i], books[i+1:]...)
				return c.SendString("Book deleted")
			}
		}
		return c.Status(404).SendString("Book not found")
	})

	// Listen on port 3000
	app.Listen(":3000")
}

func atoi(s string) int {
	var res int
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		res = res*10 + int(c-'0')
	}
	return res
}
