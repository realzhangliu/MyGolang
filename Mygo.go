package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	"encoding/json"
	"runtime"
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

func MapF(f func(x, y int) int, x1, y1 int) {
	fmt.Println(f(x1, y1))
}

type state_confirm struct {
	key  int
	done chan string
}

func main() {
	defer func() {
		fmt.Print("\nFinally!")
	}()
	//var wg sync.WaitGroup
	//
	//statful := make(chan state_confirm, 1)
	//files, err := ioutil.ReadDir(path)
	//if err == nil {
	//
	//	wg.Add(len(files))
	//
	//	for _, item := range files {
	//		fmt.Println("OPEN FILE:" + item.Name())
	//		if !item.IsDir() {
	//			go ProcessJsonContent(item, statful, &wg)
	//		}
	//	}
	//}
	//mm:=make(map[int]state_confirm)
	//for i := 0; i < len(files); i++ {
	//	mm[i]=<-statful
	//	fmt.Println(mm[i].key)
	//}
	//
	//wg.Wait()
	//ex, _ := os.Executable()
	//dir := filepath.Dir(ex)
	//list, _ := ioutil.ReadDir(dir)
	//for _,v := range list {
	//	fmt.Println(string(v.Name()))
	//}

}

func ProcessJsonContent(f os.FileInfo, sc chan state_confirm, wg *sync.WaitGroup) {
	stateful := state_confirm{key: rand.Intn(100), done: make(chan string, 1)}
	strB, _ := json.Marshal("Json String." + f.Name())
	//fmt.Println(string(strB))
	sc <- stateful
	stateful.done <- string(strB)
	time.Sleep(time.Second * 2)
	//close(stateful.done)
	defer wg.Done()
}
