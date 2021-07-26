package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	Id   string
	Name string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	name, err := queryNameById(3)
	if err != nil {
		fmt.Println("query err: ", err)
		return
	}

	fmt.Println("name: ", name)

}

func queryNameById(id int) (string, error) {
	var name string
	sqlstr := "select name from user where id = ?"
	if db == nil {
		return "no db", nil
	}

	fmt.Println(db)
	err := db.QueryRow(sqlstr, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("not find")
			return "", nil
		} else {
			fmt.Println("queryNameById error: ", err)
			return "", errors.Wrap(err, "queryNameById err")
		}
	}
	return name, nil
}
