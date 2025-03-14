package main

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID       int       `json:"id"`
    Name     string    `json:"name"`
    Age      int       `json:"age"`
    Comments []Comment `json:"comments"`
}

type Comment struct {
    ID     int    `json:"id"`
    Text   string `json:"text"`
    UserID int    `json:"user_id"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Age INTEGER NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	// Создаем таблицу комментариев
	query = `CREATE TABLE IF NOT EXISTS comments (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Text TEXT NOT NULL,
		UserID INTEGER,
		FOREIGN KEY (UserID) REFERENCES users(ID) ON DELETE CASCADE
	);`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func prepareQuery(operation string, table string, user User) interface{}{
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	switch operation {
	case "insert":
		return sq.Insert(table).Columns("Name", "Age").Values(user.Name, user.Age)
	case "select":
		return sq.Select("ID", "Name", "Age").From(table).Where(squirrel.Eq{"ID": user.ID})
	case "delete":
		return sq.Delete(table).Where(squirrel.Eq{"ID": user.ID})
	case "update":
		return sq.Update(table).Set("Name", user.Name).Set("Age", user.Age).Where(squirrel.Eq{"ID": user.ID})
	default:
		return fmt.Errorf(operation)
	}
}

func InsertUser(user User) error{
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = prepareQuery("insert", "users", user).(*squirrel.InsertBuilder).RunWith(db).Exec()
	if err != nil{
		return err
	}
	return nil
}

func SelectUser(userID int) (User, error){
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	user := User{
		ID: userID,
	}
	row, err := prepareQuery("select", "users", user).(*squirrel.SelectBuilder).RunWith(db).Query()
	if err != nil{
		return user, err
	}

	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil{
		return user, err
	}
	return user, nil
}

func UpdateUser(user User) error{
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = prepareQuery("update", "users", user).(*squirrel.UpdateBuilder).RunWith(db).Exec()
	if err != nil{
		return err
	}
	return nil
}

func DeleteUser(userID int) error{
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	user := User{
		ID: userID,
	}
	_, err = prepareQuery("delete", "users", user).(*squirrel.DeleteBuilder).RunWith(db).Exec()
	if err != nil{
		return err
	}
	return nil
}
func main(){

}