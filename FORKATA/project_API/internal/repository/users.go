package database

import (
	"context"
	"database/sql"
	"log"
)

type User struct{
	Id int
	Name string
	Email string
}

type PostgresDb struct{
	db *sql.DB
}

type Conditions struct{
	Limit int
	Offset int
}

func NewDb(db *sql.DB) *PostgresDb{
	return &PostgresDb{
		db: db,
	}
}

type UserRepository interface {
    Create(ctx context.Context, user User) error
    GetByID(ctx context.Context, id string) (User, error)
    Update(ctx context.Context, user User) error
    Delete(ctx context.Context, id string) error
	List(ctx context.Context, c Conditions) ([]User, error)
}

func (d *PostgresDb) Create(ctx context.Context, user User) error{
	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3)`
	_, err := d.db.ExecContext(ctx, query,user.Id, user.Name, user.Email)
	if err != nil{
		log.Println("Ошибка в создании пользователя")
		return err
	}
	return nil
}

func (d *PostgresDb) GetByID(ctx context.Context, id string) (User, error){
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	row := d.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil{
		log.Println("Ошибка в поиске в БД")
		return User{}, err
	}
	return user, nil
}

func (d *PostgresDb) Delete(ctx context.Context, id string) error{
	query := `DELETE FROM users WHERE id = $1 ` 
	_, err := d.db.ExecContext(ctx, query, id)
	if err != nil{
		log.Println("Ошибка в удалении пользователя")
		return err
	}
	return nil
}

func (d *PostgresDb) Update(ctx context.Context, user User) error{
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3;`
	_, err := d.db.ExecContext(ctx, query, user.Name, user.Email, user.Id)
	if err != nil{
		log.Println("Ошибка в Update")
		return err
	}
	return nil
}

func (d *PostgresDb) List(ctx context.Context, c Conditions) ([]User, error) {
    query := `SELECT id, name, email FROM users
              ORDER BY name
              LIMIT $1 OFFSET $2`

    rows, err := d.db.QueryContext(ctx, query, c.Limit, c.Offset)
    if err != nil {
        log.Println("Ошибка при выполнении List:", err)
        return nil, err
    }
	    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
            log.Println("Ошибка при сканировании пользователя:", err)
            return nil, err
        }
        users = append(users, u)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
	
    return users, nil
}