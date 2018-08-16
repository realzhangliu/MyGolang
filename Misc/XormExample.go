package Misc

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type labor struct {
	Id         int64
	Name       string `xorm:"varchar(20) notnull 'name'"`
	Age        int    `xorm:"int(3) notnull "`
	Occupation string `xorm:"varchar(30) notnull"`
}

var DBSource string = "ml:0000@tcp(192.168.1.109)/person?charset=utf8mb4"
var DBSource_loal string = "root:0000@/person?charset=utf8mb4"

func XormExample() {
	var err error
	engine, err = xorm.NewEngine("mysql", DBSource)
	checkError(err)
	err = engine.Ping()
	var lb labor
	_, err = engine.Table("labor").Where("ID=?", "1").Get(&lb)
	fmt.Println(lb)

	var pwd []string
	engine.Table("labor").Cols("name").Find(&pwd)
	fmt.Println(pwd)

	cols := []string{"name", "age"}
	var valuesSlice = make([]interface{}, len(cols))
	engine.Table("labor").Cols(cols...).Get(&valuesSlice)
	fmt.Println(string(valuesSlice[1].([]byte)))

	lb = labor{
		Name: "zl",
	}
	ok, err := engine.Table("labor").Exist(&lb)
	if ok {
		fmt.Println("FOUND.")
	}


}

func backup() {
	engine.Insert(&labor{
		Name:       "xorm",
		Age:        233,
		Occupation: "Golang package",
	})
}
