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
	case "QucikSort":
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
	SpecialName string
}

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	//ex, _ := os.Executable()
	//fmt.Println(ex)
	//dir := filepath.Dir(ex)
	//list, _ := ioutil.ReadDir(dir)
	//for _, v := range list {
	//	fmt.Println(string(v.Name()))
	//}

	//algorithm_start("MergeSort")

	for sum := 0; sum < 100; sum++ {
		fmt.Printf("Type:%T Value:%v",sum,sum)
	}

}
