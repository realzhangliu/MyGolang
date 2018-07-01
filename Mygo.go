package main

import (
	"fmt"
	"math/rand"

	_ "github.com/netldds/MyGolang/Hash"
	"encoding/base64"
	"crypto/sha1"
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

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	//RunMatrixMutiply()
	//MyUploadServer.Start()
	//Hash.NewHash()
	//DataBaseOperation.MysqlExample()
	//DataBaseOperation.Sqllite3Example()
	b := make([]byte, 32)
	rand.Read(b)
	shNew:=sha1.New()
	shNew.Write(b)
	b=shNew.Sum(nil)
	baseStr:=base64.URLEncoding.EncodeToString(b)
	fmt.Println(baseStr)
}
