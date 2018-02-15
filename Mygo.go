package main

import (
	"fmt"
	"math/rand"
	"os"
)

var mapOne = make(map[string]int)

type ctt struct {
	id    int
	state string
	rate  int
}

func schedule() {

}

func main() {
	//var a = make([]int, 0)
	defer func() {
		fmt.Print("\nFinally!")
	}()

	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	//MergeSort(&a, 0, len(a)-1)
	//QuickSort(&a, 0, len(a)-1)
	//fmt.Printf("Sorted array is :%v\n", a)

	//chh := make(chan ctt)
	//go func(ch chan ctt) {
	//	defer close(chh)
	//	time.Sleep(5*time.Second)
	//	ch<-ctt{
	//		id:    0,
	//		state: "ok",
	//		rate:  99,
	//	}
	//}(chh)
	//fmt.Println(<-(chh))
	fileInfo, err := os.Stat("tt.zip")
	if err != nil {
		fmt.Println((s.Args[0]/1<<20)
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(fileInfo.Size())

}
