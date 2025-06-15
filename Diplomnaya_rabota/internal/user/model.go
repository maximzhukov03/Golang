package user

import "time"

type User struct {
    ID        int64     // Уникальный идентификатор пользователя
    Email     string    // Email (уникальный)
    Password  string    // Хешированный пароль
    CreatedAt time.Time // Время регистрации
}