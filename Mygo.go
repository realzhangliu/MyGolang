package main

import (
	"fmt"
	"github.com/golang/glog"

	"bufio"
	"os"

	"flag"
	"reflect"
	"time"
)

const (
	file1 = "6.9MB.jpg"
	file2 = "26MB.jpg"
	file3 = "1MB.gif"
	file4 = "3.7MB.tga"
	file5 = "13.4MB.pdf"
)

var counterPool = make(map[string]time.Time)

const rootpath = "/home/dx/GoWorkBench/src/dx/taishan/data/comment_files"
const hextable = "0123456789abcdef"

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	var none interface{}
	name := "123 '123' \r dfdfd"
	none = name
	if v,ok:=none.(string);ok{
		glog.Info(v)
	}
}
func InputLoop() {
	rd := bufio.NewReader(os.Stdin)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	}

}
func reflect_example() {
	method := func(i string) map[string]string {
		fmt.Println(i)
		return map[string]string{"id": "id1"}
	}
	var v reflect.Value

	v = reflect.ValueOf(method)
	vs := []reflect.Value{reflect.ValueOf("myStrings")}
	rvs := v.Call(vs)
	res := rvs[0].MapIndex(reflect.ValueOf("id")).String()
	fmt.Println(res)
}
