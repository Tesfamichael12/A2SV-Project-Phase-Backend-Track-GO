package models

type Member struct {
	ID            int
	Name          string
	Password      string
	BorrowedBooks []Book // slice of Book
}
