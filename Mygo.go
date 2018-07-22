package main

import (
	"fmt"
	"math/rand"

	_ "github.com/netldds/MyGolang/Hash"
	"github.com/netldds/MyGolang/MyUploadServer"
)

func algorithmStart(name string) {

	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}
	switch name {
	case "QuickSort":
		QuickSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	case "MergeSort":
		MergeSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	default:
		fmt.Println("Correct name was expected.")
	}

}

type Recurlyservers struct {
	JsonName string `json:"json_name"`
}

func main() {
	defer func() {
		fmt.Print("\n Processing has done!")
	}()
	//RunMatrixMutiply()
	MyUploadServer.Start()
	//Hash.NewHash()
	//DataBaseOperation.MysqlExample()
	//DataBaseOperation.Sqllite3Example()

}
