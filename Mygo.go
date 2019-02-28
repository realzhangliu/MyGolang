package main

import (
	"fmt"

	"bufio"
	"os"

	"reflect"
)

type Order struct {
	OrdId      int
	customerId int
}
type GG interface {
	gg() string
}

func (order *Order) gg() string {
	return "gg"
}
func main() {
	g1 := &Order{}
	fmt.Println(g1.gg())
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
