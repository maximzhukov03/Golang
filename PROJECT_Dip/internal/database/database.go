package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
)

// DB глобальная переменная для доступа к БД
var DB *gorm.DB

// InitDB инициализирует подключение к БД
func InitDB() (*gorm.DB, error) {
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "./storage.db" // значение по умолчанию
    }

    db, err := gorm.Open(sqlite.Open("file:"+os.Getenv("DB_PATH")+"?cache=shared&_journal_mode=WAL"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    DB = db
    return db, nil
}

// Migrate выполняет миграции базы данных
func Migrate(db *gorm.DB) error {
    // Здесь будем регистрировать модели для миграции
    err := db.AutoMigrate(
        // Добавьте здесь свои модели
        // &model.User{},
        // &model.File{},
    )
    return err
}