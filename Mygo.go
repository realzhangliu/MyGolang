package main

import (
	"fmt"
	"math/rand"
)

func algorithm_start(name string) {

	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	switch name {
	case "QucikSort":
		QuickSort(&a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	case "MergeSort":
		MergeSort(&a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	default:
		fmt.Println("Correct name was expected.")
	}

}

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	if err:=test_t_interface();err!=nil{
		fmt.Println(err.Check())
	}
}
