package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// 检查错误
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 事务错误
func checkTxErr(err error, tx *sql.Tx) {
	if err != nil {
		log.Println(err)
		err = tx.Rollback()
		checkErr(err)
	}
}

// 数据库实体
type User struct {
	ID      int
	Name    string
	Age     int
	Sex     int
	AddDate time.Time
}

func main() {
	// 1、获取数据库连接
	conn := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "test",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	// fmt.Println(conn.FormatDSN())
	db, err := sql.Open("mysql", conn.FormatDSN())
	checkErr(err)
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	// err = db.Ping()
	// checkErr(err)

	if err := db.Ping(); err != nil {
		fmt.Println("db.Ping", err)
		return
	}

	// 3、创建表
	// sql := `
	// 	CREATE TABLE IF NOT EXISTS users(
	//         id INT NOT NULL AUTO_INCREMENT,
	//         name VARCHAR(100) NOT NULL,
	//         age INT NOT NULL,
	//         sex TINYINT,
	//         add_date DATETIME,
	//         PRIMARY KEY(id)
	//     )
	// `

	// _, err = db.Exec(sql)
	// checkErr(err)

	// 4、添加数据/
	// sql := "INSERT INTO users (name,age,sex,add_date) VALUES (?,?,?,?)"
	// res, err := db.Exec(sql, "Jerry", 18, 1, time.Now())
	// checkErr(err)
	// fmt.Println(res.LastInsertId())

	// 5、查询数据
	// sql := "SELECT id,name,age,sex,add_date FROM users"
	// rows, err := db.Query(sql)
	// checkErr(err)
	// defer rows.Close()
	// user := User{}
	// for rows.Next() {
	// 	err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	// 	checkErr(err)
	// 	fmt.Println(user, user.AddDate.Format("2022/04/01 16:50:05"))
	// }

	// err = rows.Err()
	// checkErr(err)

	// 6、查询一行
	// sql := "SELECT id,name,age,sex,add_date FROM users"
	// row := db.QueryRow(sql)
	// user := User{}
	// err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	// checkErr(err)
	// fmt.Println(user)

	// 7、命令
	// sql := "UPDATE users SET name=? WHERE id=?"
	// stmt, err := db.Prepare(sql)
	// checkErr(err)
	// defer stmt.Close()
	// result, err := stmt.Exec("adiachan", 1)
	// checkErr(err)
	// fmt.Println(result.RowsAffected())

	// 8、查询
	// sql := "SELECT id,name,age,sex,add_date FROM users WHERE id=?"
	// stmt, err := db.Prepare(sql)
	// checkErr(err)
	// defer stmt.Close()
	// row := stmt.QueryRow(1)
	// user := User{}
	// err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	// checkErr(err)
	// fmt.Println(user)

	// 9、事务 transaction
	tx, err := db.Begin()
	checkErr(err)
	_, err = tx.Exec("UPDATE users SET age=? WHERE id=?", 20, 1)
	checkTxErr(err, tx)
	_, err = tx.Exec("update users SET sex=? WHERE id=?", 0, 1)
	checkTxErr(err, tx)
	err = tx.Commit()
	checkTxErr(err, tx)

	// 10、查询
	// sql := "SELECT id,name,age,sex,add_date FROM users WHERE id=?"
	// row := db.QueryRow(sql, 3)
	// user := User{}
	// err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Sex, &user.AddDate)
	// checkErr(err)
	// fmt.Println(user)
}
