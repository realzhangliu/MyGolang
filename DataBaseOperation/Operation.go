package DataBaseOperation

import (
	"database/sql"
	"fmt"
	"log"
	//仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
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
