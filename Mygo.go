package main

import (
	"fmt"
	"math/rand"

	"bufio"
	"os"

	"reflect"

	_ "github.com/netldds/MyGolang/Hash"
	"github.com/netldds/MyGolang/Misc"
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

type b struct {
	counter int
}

type hibt interface {
	binterface()
}
type otherbt interface {
	ointerface()
}

func (e b) binterface() {
	fmt.Println("print b interface.")
	fmt.Println(e.counter)
}
func (e b) ointerface() {
	fmt.Println("print o interface.")
	fmt.Println(e.counter)
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
	Misc.XormExample()
	//Misc.TimeExample()
	//var i []int=[]int{1,2,3}

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
