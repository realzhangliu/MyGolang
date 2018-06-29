package DataBaseOperation

import (
	"database/sql"
	"fmt"
	"log"
	//仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
	_ "github.com/go-sql-driver/mysql"
)

//user@unix(/path/to/socket)/dbname?charset=utf8
//user:password@tcp(localhost:5555)/dbname?charset=utf8
//user:password@/dbname
//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

func Man() {
	db, err := sql.Open("mysql", "root:0000@/account?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	stmt, err := db.Prepare("INSERT userinfo set username=?,department=?,created=?")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	res, err := stmt.Exec("one", "one-dept", "2018.6.26")
	res, err = stmt.Exec("two", "one-dept", "2018.6.26")
	res, err = stmt.Exec("three", "one-dept", "2018.6.26")
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

	stmt, err = db.Prepare("UPDATE userinfo set username=? WHERE uid=?")
	checkErr(err)
	res, err = stmt.Exec("666", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("SELECT * from userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	stmt, err = db.Prepare("DELETE  from userinfo")
	checkErr(err)
	res, err = stmt.Exec()
	checkErr(err)
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
