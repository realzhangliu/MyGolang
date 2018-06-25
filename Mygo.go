package main

import (
	"fmt"
	"math/rand"

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

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	b:=[]int{1,2,3,4,5}
	for v:=range b{
		fmt.Println(v)
	}
	RunMatrixMutiply()
	MyUploadServer.Start()

}
