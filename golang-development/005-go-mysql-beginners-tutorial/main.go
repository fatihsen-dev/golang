package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySql Tutorial")

	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/go_mysql_tutorial_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Successfully inserted into user tables")

	results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
