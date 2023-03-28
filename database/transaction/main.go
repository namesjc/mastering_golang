package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	Age       int
	Sex       int
	Create_at time.Time
}

func main() {
	db, err := newDB()
	if err != nil {
		fmt.Println("newDB", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	newName := func(s string) *string {
		return &s
	}

	newAge := func(b int) *int {
		return &b
	}

	newSex := func(c int) *int {
		return &c
	}

	if err := insertUsers(db, []User{
		{
			Name: *newName("test3"),
			Age:  *newAge(10),
			Sex:  *newSex(0),
		},
		{
			Name: *newName("test4"),
			Age:  *newAge(11),
			Sex:  *newSex(1),
		},
	}); err != nil {
		fmt.Println("inserUsers", err)
	}

}

func insertUsers(db *sql.DB, users []User) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	for _, user := range users {
		_, err := tx.ExecContext(context.Background(), "INSERT INTO users (name, age, sex) VALUES ($1, $2,$3)", user.Name, user.Age, user.Sex)
		if err != nil {
			return fmt.Errorf("tx.ExecContex %w", err)
		}
	}

	// row := tx.QueryRow("SELECT * FROM users WHERE name='test4'")
	// user := User{}
	// if err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.Create_at); err != nil {
	// 	fmt.Errorf("row.scan %w", err)
	// }
	// fmt.Println("user", user)

	// _, err = tx.ExecContext(context.Background(), "UPDATE users SET age=age+$1 WHERE name='John'", 11)
	// if err != nil {
	// 	fmt.Errorf("tx.ExecContext %w", err)
	// }

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx.Commit %w", err)
	}
	return nil
}

func newDB() (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgresql",
		Host:   "127.0.0.1:5432",
		User:   url.UserPassword("root", "password"),
		Path:   "test",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()
	// fmt.Println("dsn.String", dsn.String())

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	return db, nil
}
