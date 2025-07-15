package main

import (
	"LibraryManagementSystem/controllers"
	"LibraryManagementSystem/models"
	"LibraryManagementSystem/services"
	"fmt"
)

func main() {
	library := &services.Library{
		Members: make(map[int]*models.Member),
		Books:   make(map[int]*models.Book),
	}

	var currentMember *models.Member
	var nextMemberID int = 1

	fmt.Println("Welcome to the Library Management System!")
	fmt.Print("Please create an account.\nEnter your username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	var password string
	fmt.Scanln(&password)

	currentMember = &models.Member{ID: nextMemberID, Name: username, Password: password, BorrowedBooks: []models.Book{}}
	library.Members[nextMemberID] = currentMember
	fmt.Printf("Account created! Your Member ID is %d.\n\n", nextMemberID)
	nextMemberID++

	for {
		fmt.Printf(`
What service would you like to get?
(Enter the number associated with the service you want)
1. Add a Book
2. Remove a Book
3. Borrow a Book
4. Return a Book
5. List Available Books
6. List Borrowed Books
`)
		var serviceNo int
		_, err := fmt.Scanln(&serviceNo)
		if err != nil {
			fmt.Println("Wrong Input. Please Try Again!")
			continue
		}
		switch serviceNo {
		case 1:
			controllers.HandleAddBook(library)
		case 2:
			controllers.HandleRemoveBook(library)
		case 3:
			controllers.HandleBorrowBook(library)
		case 4:
			controllers.HandleReturnBook(library)
		case 5:
			controllers.HandleListAvailableBooks(library)
		case 6:
			controllers.HandleListBorrowedBooks(library)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}