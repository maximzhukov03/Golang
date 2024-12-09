package postgres

import (
	"database/sql"
	"golandg/sql/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository { //создание струткуры баззы данных
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]domain.User, error) { // Метод для получения всех юзеров
	rows, err := r.db.Query("SELECT * FROM user_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *userRepository) GetByID(id int) ([]domain.User, error) {
	rows, err := r.db.Query("SELECT * FROM user_data where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.SecondName, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *userRepository) Create(user domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO user_data (first_name, second_name, email, password) VALUES ($1, $2, $3, $4)",
		user.FirstName, user.SecondName, user.Email, user.Password,
	)
	return err
}

func (r *userRepository) Delete(user domain.User) error {
	_, err := r.db.Exec("DELETE FROM user_data where first_name = $1", user.FirstName)
	return err
}
