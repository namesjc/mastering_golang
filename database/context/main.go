package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	ID        int
	Name      string
	Age       int
	Sex       int
	Create_at time.Time
}

func main() {
	dsn := url.URL{
		Scheme: "postgresql",
		Host:   "127.0.0.1:5432",
		User:   url.UserPassword("root", "password"),
		Path:   "test",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()
	fmt.Println("dsn.String", dsn.String())

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
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

	res, err := db.ExecContext(context.Background(), "INSERT INTO users (name,age,sex) VALUES ('Echo', 18, 1)")
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}
	fmt.Println(res.LastInsertId())

	row := db.QueryRowContext(context.Background(), "SELECT * FROM users WHERE name='John'")
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	user := User{}

	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.Create_at); err != nil {
		fmt.Println("row.scan", err)
		return
	}

	fmt.Println("user", user)

	rows, err := db.QueryContext(context.Background(), "SELECT id, name, age, sex, create_at FROM users")
	if err != nil {
		fmt.Println("rwo.Scan", err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		fmt.Println("row.Err()", err)
		return
	}

	for rows.Next() {
		user = User{}

		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.Create_at); err != nil {
			fmt.Println("rows.Scan", err)
			return
		}

		fmt.Println(user)
	}
}
