package models

import (
	"fmt"

	"github.com/sangmin4208/bookstore-management-api-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}
func (b *Book) UpdateBook() (*Book, error) {
	fmt.Println(*b)
	result := db.Save(b)
	if result.Error != nil {
		return nil, result.Error
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	if result := db.Find(&Books); result.Error != nil {
		return nil, result.Error
	}
	return Books, nil
}

func GetBookById(id uint) (*Book, error) {
	var Book Book
	if result := db.First(&Book, id); result.Error != nil {
		return nil, result.Error
	}
	return &Book, nil
}

func DeleteBook(ID uint) (*Book, error) {
	var Book Book
	if result := db.First(&Book, ID); result.Error != nil {
		return nil, result.Error
	}
	if result := db.Delete(&Book); result.Error != nil {
		return nil, result.Error
	}
	return &Book, nil
}
