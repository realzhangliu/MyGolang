package main

import (
	"flag"
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
	//PingBaidu()
	var a1,b1,c1 string
	flag.StringVar(&a1,"Alpha","Value-A","usage:Value-A")
	b1=flag.Arg(0)
	c1=flag.Arg(1)
	flag.Parse()
	fmt.Printf("%v  %v  %v  ",flag.Args(),b1,c1)
}
