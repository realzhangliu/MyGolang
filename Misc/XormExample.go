package Misc

import (
	"fmt"

	"encoding/binary"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math"
)

var engine *xorm.Engine

func XormExample() {
	var err error
	engine, err = xorm.NewEngine("mysql", "ml:0000@/person?charset=utf8mb4")
	checkError(err)
	err = engine.Ping()
	result, err := engine.Query("select * from labor")
	for _, v := range result {
		id := binary.LittleEndian.Uint32(v["ID"])
		fmt.Printf("%d\n%s\n%s\n%s\n", id, string(v["name"]), v["age"], string(v["occupation"]))

	}

}
var pl=fmt.Println
func Byte2Int() {
	//var i uint32
	var b []byte
	b=make([]byte,4)
	pl(math.Float32bits(10))
	binary.LittleEndian.PutUint32(b,math.Float32bits(10))
	pl(b)

	//conversion
	bits:=binary.LittleEndian.Uint32(b)
	result:=math.Float32frombits(bits)
	pl(result)

}