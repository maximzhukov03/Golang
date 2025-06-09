package models

import()

type Pet struct {
    ID     int64  `json:"id"`
    Name   string `json:"name"`
    Status string `json:"status"`
}

