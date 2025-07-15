package controllers

import (
	"LibraryManagementSystem/models"
	"fmt"
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("library.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open log file:", err)
		return
	}
	log.SetOutput(f)
}

func HandleAddBook(library models.LibraryManager) {
	var id int
	var title, author string
	fmt.Print("Enter book ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)
	if id < 1 || title == "" || author == "" {
		fmt.Println("Error: Invalid Input")
		log.Println("Error: Invalid Input")
		return
	}
	book := models.Book{ID: id, Title: title, Author: author, Status: models.StatusAvailable}
	err := library.AddBook(book)
	if err != nil {
		fmt.Println("Error:", err)
		log.Println("Error:", err)
	} else {
		fmt.Println("Book added successfully!")
	}
}

func HandleRemoveBook(library models.LibraryManager) {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scanln(&id)
	err := library.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
		log.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func HandleBorrowBook(library models.LibraryManager) {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		log.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func HandleReturnBook(library models.LibraryManager) {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
		log.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func HandleListAvailableBooks(library models.LibraryManager) {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func HandleListBorrowedBooks(library models.LibraryManager) {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	books := library.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
