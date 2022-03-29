package book

import (
	// "errors"
	// "fmt"

	"gorm.io/gorm"
)

//We will use gorm
type BookRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *BookRepository = nil

//initilaze the repo with gorm db
func BookRepoInit(db *gorm.DB) *BookRepository {

	if singleton == nil {
		singleton = &BookRepository{db}
	}
	return singleton
}

//retunr same gorm db instance and we dont need pass it as parameter
//maybe there is a better way to handle this case
func Repo() *BookRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *BookRepository) Migrations() {
	c.db.AutoMigrate(&Book{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

//Used to Insert data to SQL
//Json to SQL :)
func (c *BookRepository) InsertSampleData(books Books) {

	for _, book := range books {
		c.db.Where(Book{ID: book.ID}).
			Attrs(Book{ID: book.ID, BookName: book.BookName}).
			FirstOrCreate(&book)
	}

}

type BookAndAuthor struct {
	ID       int
	BookName string
	Name     string
	AuthorID int
}

func (b *BookRepository) GetBookByIdIncludeAuthor(id string) []BookAndAuthor {

	var model []BookAndAuthor
	b.db.
		// First(&model).
		Joins("left join authors on authors.author_id = books.author_id").
		Where("books.id = ?", id).
		Table("books").
		Select("books.id ,books.book_name, authors.name, authors.author_id").
		Scan(&model)
	return model
}

func (b *BookRepository) GetBooksWithAuthors() (Books, error) {
	var books Books
	result := b.db.Preload("Author").Find(&books)
	// x:= b.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return b.db.Preload("Authors").Find(&books)
	// })

	// fmt.Println(x)

	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

//Just type full book name
func (c *BookRepository) FindByName(bookName string) (*Book, error) {
	// var test struct{
	// 	b
	// }
	var book *Book
	//lke quert
	result := c.db.First(&book, "Book_name like ?", "%"+ bookName +"%")
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (c *BookRepository) GetByID(bookID string) (*Book, error) {
	var book *Book
	result := c.db.First(&book, "ID = ?", bookID).Find(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (b *BookRepository) UpdateBookQuantity(bookID, quantity string) error {
	var book *Book
	// result := b.db.Update(&book, "stock_count", quantity).Where("id = ?", bookID)
	result := b.db.Model(&book).Where("ID = ?", bookID).Update("stock_count", quantity)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
func (b *BookRepository) SoftDeleteBook(bookID string) error {
	result := b.db.Where("id = ?", bookID).Delete(&Book{})

	//db.Where("age = ?", 20).Delete(&User{})

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (b *BookRepository) DropTable() error {
	result := b.db.Exec("DROP TABLE books")

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (r *BookRepository) GetAllBooks() (Books, error) {
	var books Books
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// func (c *BookRepository) FindBooks(bookID string) (*Books, error) {
// 	var book *Books
// 	result := c.db.Find(&book,)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return book, nil
// }

// func (c *BookRepository) UpdateBookQuantity(bookName string, quantity int) (*Book, error) {
// 	var book *Book
// 	result := c.db.Update(&book, "Book_name = ?", bookName)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return book, nil
// }
