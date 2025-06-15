package models

import ()

type Author struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Popularity int `json:"popularity"`
    Books []Book `json:"books",omitempty`
}