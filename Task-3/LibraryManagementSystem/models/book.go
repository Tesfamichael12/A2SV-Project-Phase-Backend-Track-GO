package models

import (
	"errors"
)


const (
	StatusAvailable = "Available"
	StatusBorrowed  = "Borrowed"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status string 
}

func (b *Book) SetStatus(s string) error {
	if s != StatusAvailable && s != StatusBorrowed {
		return errors.New("invalid status value")
	}
	b.Status = s
	return nil
}

