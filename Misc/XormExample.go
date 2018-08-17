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

var pln = fmt.Println

func XormExample() {
	var err error
	engine, err = xorm.NewEngine("mysql", DBSource_loal)
	checkError(err)
	err = engine.Ping()

	//Query one record
	pln("Query one record")
	var lb labor
	_, err = engine.Table("labor").Where("ID=?", "1").Get(&lb)
	pln(lb)

	//Query multiple records
	pln("Query multiple records")
	var pwd []string
	engine.Table("labor").Cols("name").Find(&pwd)
	pln(pwd)

	//Query one record
	pln("Query one record")
	cols := []string{"name", "age"}
	var valuesSlice = make([]interface{}, len(cols))
	engine.Table("labor").Cols(cols...).Get(&valuesSlice)
	pln(string(valuesSlice[1].([]byte)))

	//Insert
	pln("Insert")
	pln(lb)
	lb.Id = 0
	//labor{Name: "tt",
	//	Age:        11,
	//	Occupation: "programmer"}
	//i64, _ := engine.Insert(lb)
	//pln(i64)

	//check if one record is exist on table.
	pln("check if one record is exist on table.")
	lb = labor{
		Name: "zl",
	}
	ok, err := engine.Table("labor").Exist(&lb)
	if ok {
		pln("FOUND.")
	}
	if ok, err = engine.Table("labor").Where("id=?", 2).Exist(); ok {
		pln("id=2 found.")
	}

	//query multiple records
	pln("query multiple records")
	var lbs []labor
	err = engine.Table("labor").Where("name=?", "zl").And("age >?", 1).Limit(10, 0).Find(&lbs)
	pln(lbs)

	var ress []struct {
		Id             int    `xorm:"id"`
		Name           string `xorm:"name"`
		Age            int    `xorm:"age"`
		Occupation     string `xorm:"occupation"`
		Id_bak         int    `xorm:"id_bak"`
		Name_bak       string `xorm:"name_bak"`
		Occupation_bak string `xorm:"occupation_bak"`
		Password_bak   string `oxrm:"password_bak"`
	}
	err = engine.Table("labor").Select("labor.ID,labor.name,labor.age,labor.occupation,labor_bak.ID as 'id_bak',labor_bak.name as 'name_bak',labor_bak.occupation as 'occupation_bak',labor_bak.password as 'password_bak'").Join("INNER", "labor_bak", "labor.age=labor_bak.age").Limit(10, 0).Where("labor.age=?", 28).Find(&ress)
	pln(ress)

}

func backup() {
	engine.Insert(&labor{
		Name:       "xorm",
		Age:        233,
		Occupation: "Golang package",
	})
}
