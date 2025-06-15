package models

import (
    "time"
)

// User представляет пользователя системы
// Поле PasswordHash хранится лишь в БД и не возвращается в API
// Поле CreatedAt устанавливается при создании в репозитории
// JSON-теги используются для сериализации в HTTP-ответах
// DB-теги можно использовать, если вы используете sqlx или другую ORM

type User struct {
    ID           int64     `db:"id" json:"id"`
    Email        string    `db:"email" json:"email"`
    PasswordHash string    `db:"password_hash" json:"-"`
    CreatedAt    time.Time `db:"created_at" json:"created_at"`
}


// File представляет загруженный файл с метаданными и ссылкой

type File struct {
    ID         int64     `db:"id" json:"id"`
    UserID     int64     `db:"user_id" json:"user_id"`
    Name       string    `db:"name" json:"name"`
    Size       int64     `db:"size" json:"size"`
    Bucket     string    `db:"bucket" json:"bucket"`
    ObjectName string    `db:"object_name" json:"object_name"`
    UploadedAt time.Time `db:"uploaded_at" json:"uploaded_at"`
    // URL формируется в слое handler/service, поэтому не сохраняется в БД
    URL        string    `db:"-" json:"url,omitempty"`
}
