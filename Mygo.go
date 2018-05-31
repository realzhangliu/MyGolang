package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
)

func algorithm_start(name string) {

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
	//MyUploadServer.Start()

	h := md5.New()
	io.WriteString(h, "hello")
	fmt.Printf("%x", h.Sum(nil))

}
