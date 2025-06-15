package models

import()

type Book struct {
    ID int64 `json:"id"`
    Title string `json:"title"`
    AuthorID string `json:"author_id"`
    UserID string `json:"user_id"`
}