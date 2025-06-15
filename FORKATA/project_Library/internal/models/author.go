package models

import ()

type Author struct {
    ID       int64     `json:"id"`
    Name string `json:"name" db:"name"`
    Books []Book `json:"books",omitempty`
}