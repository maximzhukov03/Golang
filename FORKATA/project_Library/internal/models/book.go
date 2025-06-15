package models

import()

type Book struct {
    ID string `json:"id"`
    Title string `json:"title"`
    AuthorID string `json:"author_id"`
    UserID string `json:"user_id"`
}