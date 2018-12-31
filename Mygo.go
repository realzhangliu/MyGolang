package main

import (
	"fmt"

	"bufio"
	"os"

	"reflect"

	"github.com/tus/tusd"
)

type Order struct {
	OrdId      int
	customerId int
}

func main() {

	v := make(map[string]interface{})
	v["a"] = func() {
		fmt.Println("lalala")
	}
	v["b"] = 2
	vr := reflect.ValueOf(&v)
	fmt.Println(vr.Elem().MapIndex(vr.Elem().MapKeys()[1]).Elem())
	fmt.Println(vr.Elem().MapIndex(vr.Elem().MapKeys()[1]))

	handler ,_:= tusd.NewHandler(tusd.Config{})

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
	var inf interface{}

	inf = method
	v = reflect.ValueOf(inf)
	vs := []reflect.Value{reflect.ValueOf(interface{}("myStrings"))}
	rvs := v.Call(vs)
	res := rvs[0].MapIndex(reflect.ValueOf("id")).String()
	fmt.Println(res)
}
