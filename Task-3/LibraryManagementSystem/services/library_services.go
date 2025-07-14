package services

import (
	"LibraryManagementSystem/models"
	"errors"
)

type Library struct {
	Members map[int]*models.Member
	Books   map[int]*models.Book
}


func (l *Library) AddBook(book models.Book) error {
	if _, exists := l.Books[book.ID]; exists {
		return errors.New("book already exists")
	}
	book.Status = models.StatusAvailable
	l.Books[book.ID] = &book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book does not exist")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member does not exist")
	}
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status != models.StatusAvailable {
		return errors.New("book is not available")
	}
	err := book.SetStatus(models.StatusBorrowed)
	if err != nil {
		return err
	}
	member.BorrowedBooks = append(member.BorrowedBooks, *book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member does not exist")
	}
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book does not exist")
	}
	if book.Status != models.StatusBorrowed {
		return errors.New("book was not borrowed")
	}
	err := book.SetStatus(models.StatusAvailable)
	if err != nil {
		return err
	}
	// Remove book from member's BorrowedBooks slice
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)
	for _, book := range l.Books {
		if book.Status == models.StatusAvailable {
			availableBooks = append(availableBooks, *book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}