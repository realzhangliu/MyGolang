package Misc

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var engine *xorm.Engine

type labor struct {
	Id         int64
	Name       string `xorm:"varchar(20) notnull 'name'"`
	Age        int    `xorm:"int(3) notnull "`
	Occupation string `xorm:"varchar(30) notnull"`
}

func XormExample() {
	var err error
	engine, err = xorm.NewEngine("mysql", "ml:0000@/person?charset=utf8mb4")
	checkError(err)
	err = engine.Ping()
	var lb labor
	_, err = engine.Table("labor_bak").Where("ID=?", "2").Get(&lb)
	//fmt.Println(lb, flag)

	var pwd []string
	engine.Table("labor_bak").Cols("password").Find(&pwd)
	//fmt.Println(pwd)

	cols := []string{"name", "age", "password"}
	var valuesSlice = make([]interface{}, len(cols))
	engine.Table("labor_bak").Cols(cols...).Get(&valuesSlice)
	fmt.Println(string(valuesSlice[1].([]byte)))
time.Now().Format()
}

func backup() {
	engine.Insert(&labor{
		Name:       "xorm",
		Age:        233,
		Occupation: "Golang package",
	})
}
