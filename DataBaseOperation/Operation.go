package DataBaseOperation

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Man() {
	db, err := sql.Open("mysql", "root:0000@/world?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	stmt, err := db.Prepare("INSERT city set Name=?")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	res, err := stmt.Exec("sixsixsix")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(id)
}
