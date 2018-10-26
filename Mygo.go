package main

import (
	"bufio"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"reflect"
	"math/rand"
	"MyGolang/Misc"
	"MyGolang/DataBaseOperation"
)



const (
	file1 = "6.9MB.jpg"
	file2 = "26MB.jpg"
	file3 = "1MB.gif"
	file4 = "3.7MB.tga"
	file5 = "13.4MB.pdf"
)
func main() {
	//DataBaseOperation.RungOrm()
	DataBaseOperation.RunnerRedis()

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
func algorithmStart(name string) {

	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}
	switch name {
	case "QuickSort":
		Misc.QuickSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	case "MergeSort":
		Misc.MergeSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	default:
		fmt.Println("Correct name was expected.")
	}

}