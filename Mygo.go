package main

import (
	"fmt"
	"math/rand"

	"github.com/netldds/MyGolang/Matrix"
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
	// MyUploadServer.Start()
	mt := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
		[]float32{7, 8, 9}}
	mv := [][]float32{
		{1, 2},
		{1, 2},
		{1, 2}}
	mt1, _ := matrix.New(mt)
	mt2, _ := matrix.New(mv)
	result, _ := mt1.Multiply(mt2)
	fmt.Println(result)
}
