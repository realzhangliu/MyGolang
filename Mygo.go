package main

import (
	"fmt"
	"math/rand"
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

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	//RunMatrixMutiply()
	//MyUploadServer.Start()
	//Hash.NewHash()
	//DataBaseOperation.Man()
	createQuery(nil)
}
