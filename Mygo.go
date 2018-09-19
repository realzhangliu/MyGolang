package main

import (
	"fmt"
	"math/rand"

	"bufio"
	_ "net/http/pprof"
	"os"

	"MyGolang/Misc"
	"reflect"
	"MyGolang/DataBaseOperation"
)

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

type Recurlyservers struct {
	JsonName string `json:"json_name"`
}
type II struct {
	index int
}

func calc() func() func() func() {
	fmt.Println("connect")
	return func() func() func() {
		fmt.Println("disconnect")
		return func() func() {
			fmt.Println("hahaha")
			return func() {
				fmt.Println("lalala")
			}
		}
	}
}

func main() {
	defer func() {
		fmt.Print("\n Processing has done!")
	}()
	//RunMatrixMutiply()
	//MyUploadServer.Start()
	//Hash.NewHash()
	//DataBaseOperation.MysqlExample()
	//DataBaseOperation.Sqllite3Example()
	//AtomicCounterExample()
	//MutexSample()
	//Misc.SignalExample()
	//Misc.XormExample()
	//Misc.TimeExample()
	//var i []int=[]int{1,2,3}
	//Misc.RwMutexExample()
	//http.ListenAndServe(":80", nil)
	//AuthenticationServer.StartJWT()
	DataBaseOperation.RungOrm()


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
