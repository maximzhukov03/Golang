package models

import()

type Book struct {
    ID int64 `json:"id"`
    Title string `json:"title"`
    AuthorID int `json:"author_id"`
    UserID int `json:"user_id"`
}