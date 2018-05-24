package main

import (
	"fmt"
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

type It3 struct {
	ID          int32
	Name        string
	SpecifiedName string
	promotedStruct
}
type promotedStruct struct{
	fullname string
}
type find interface{
	findsomething()
}

func (e *promotedStruct) findsomething() {
	fmt.Println("find some thing...")
}

func acceptArguments(f find)  {
		f.findsomething()
}

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	var v1 promotedStruct
	v1.fullname="zzzz"
	acceptArguments(&v1)
}

