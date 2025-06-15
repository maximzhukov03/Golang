package models

import()

type User struct {
    ID          string    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    RentedBooks []Book `json:"rented_books,omitempty"`
}