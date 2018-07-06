package main

import (
	"fmt"
	"math/rand"

	_ "github.com/netldds/MyGolang/Hash"
	"regexp"
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
		fmt.Print("\nFinally!")
	}()
	//RunMatrixMutiply()
	//MyUploadServer.Start()
	//Hash.NewHash()
	//DataBaseOperation.MysqlExample()
	//DataBaseOperation.Sqllite3Example()
	re,_:=regexp.Compile("[0-9a-zA-Z|-|\\.]*")
	result:=re.FindAllIndex([]byte("https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.3.md"),5)
	fmt.Println(result)
}
