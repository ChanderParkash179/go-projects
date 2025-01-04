package models

import (
	"bookstore/packages/config"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Book struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	Author       string    `gorm:"column:author" json:"author"`
	Publications string    `gorm:"column:publications" json:"publications"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func Init() {
	// Initialize the DB connection
	config.Connection()
	// Using the global db variable
	db = config.GetDB()

	// Ensure db is not nil
	if db == nil {
		fmt.Println("Error: DB connection failed")
		return
	}

	// AutoMigrate for Book model
	err := db.AutoMigrate(&Book{})
	if err != nil {
		// Log error if migration fails
		fmt.Println("Error migrating database:", err)
		return
	}
}

// Operations

// CreateBook creates a new book in the database
func (b *Book) CreateBook() (*Book, error) {

	response := db.Create(b)
	if response.Error != nil {
		return nil, fmt.Errorf("%s\n", response.Error)
	}

	return b, nil
}

// GetBooks returns a list of all books
func GetBooks() (*[]Book, error) {
	var books []Book

	// Passing a pointer to books
	response := db.Find(&books)

	if response.Error != nil {
		return nil, fmt.Errorf("%s\n", response.Error)
	}

	return &books, nil
}

// GetBook returns a specific book by id
func GetBook(id int64) (*Book, error) {
	var book Book
	response := db.Where("id=?", id).Find(&book)

	if response.Error != nil {
		return nil, fmt.Errorf("%s\n", response.Error)
	}

	return &book, nil
}

// DeleteBook deletes a book by id
func DeleteBook(id int64) (*Book, error) {
	var book Book
	response := db.Where("id = ?", id).Delete(&book)

	if response.Error != nil {
		return nil, fmt.Errorf("%s\n", response.Error)
	}

	return &book, nil
}
