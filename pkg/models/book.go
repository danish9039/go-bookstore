package models
//models are a way to interact witth data base with out using sql query
import (
	
    "gorm.io/gorm"
	

	"github.com/danish9039/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {

gorm.Model

Name string `gorm:""json:"name"`
Author string `json:"author"`
Publication string `json:"publication"`


}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook()*Book {

	if b.ID == 0 {
        db.Create(b) // Create a new record
    }
    return b
}


func GetAllBooks () []Book{

	var  Books []Book
	db.Find(&Books)
	return Books


	 
}

func GetBookById(Id int64) (*Book, *gorm.DB){

var getbook Book 
db:=db.Where("ID=?", Id).Find(&getbook)


 return &getbook , db


}


func DeleteBookById(Id int64) Book {

	var book Book
	db.Where( "ID=?", Id ).Delete(&book) 
	return book
}