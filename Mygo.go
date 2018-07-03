package main

import (
	"fmt"
	"math/rand"

	"time"

	_ "github.com/netldds/MyGolang/Hash"
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
	//c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		//c1 <- "string c1"
	}()
	go func() {
		time.AfterFunc(time.Second*5, func() {
			c2 <- "string c2"
		})
	}()
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	}
}
