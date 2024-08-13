package modules

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/nesuh/Book_Managment/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	var err error

	// Connect to the database
	config.Connect()
	db = config.GetDB()

	// Check if the connection was successful
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Auto-migrate the Book struct
	err = db.AutoMigrate(&Book{}).Error
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}


	func (b *Book) CreateBook() *Book{
		db.NewRecord(b)
		db.Create(&b)
		return b
	}
	func GetAllBooks() []Book{
		var Books []Book
		db.Find(&Books)
		return Books
	}
	func getBookById(Id int64)(*Book, *gorm.DB){
		var getBook Book
		db:=db.where("ID=?",Id ).Find(&getBook)
		return &getBook,db
	}

	func DeleteBook(ID int64) Book{
		var book Book
		db.Where("ID=?",ID).Delete(book)
		return book
	}
}
