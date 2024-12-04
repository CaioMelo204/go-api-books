package models

import (
	"go-api-1/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	d := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, d
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
