package main

import (
	"MyGolang/Misc"
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net"
	"os"
	"reflect"
	"time"
)

const (
	file1 = "6.9MB.jpg"
	file2 = "26MB.jpg"
	file3 = "1MB.gif"
	file4 = "3.7MB.tga"
	file5 = "13.4MB.pdf"
)

var counterPool = make(map[string]time.Time)

const rootpath = "/home/dx/GoWorkBench/src/dx/taishan/data/comment_files"
const hextable = "0123456789abcdef"

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	strs := "20190101"
	h := sha1.New()
	h.Write([]byte(strs))
	glog.Info(h.Sum(nil))
	glog.Info(hex.EncodeToString(h.Sum(nil)))

	res := make([]byte, len([]byte(h.Sum(nil)))*2)
	for i, v := range h.Sum(nil) {
		res[i*2] = hextable[v>>4]
		res[i*2+1] = hextable[v&0x0f]
	}
	glog.Info(string(res))
}

func InputLoop() {
	rd := bufio.NewReader(os.Stdin)
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(str)
	}

}
func reflect_example() {
	method := func(i string) map[string]string {
		fmt.Println(i)
		return map[string]string{"id": "id1"}
	}
	var v reflect.Value

	v = reflect.ValueOf(method)
	vs := []reflect.Value{reflect.ValueOf("myStrings")}
	rvs := v.Call(vs)
	res := rvs[0].MapIndex(reflect.ValueOf("id")).String()
	fmt.Println(res)
}
func algorithmStart(name string) {

	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(100))
	}
	switch name {
	case "QuickSort":
		Misc.QuickSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	case "MergeSort":
		Misc.MergeSort(a, 0, len(a)-1)
		fmt.Printf("Sorted array is :%v\n", a)

	default:
		fmt.Println("Correct name was expected.")
	}

}
func tcp_test() {
	//server
	go func() {
		net, _ := net.Listen("tcp", ":6666")
		fmt.Println(net.Addr())
		for {
			conn, err := net.Accept()
			fmt.Println(" server conn accepted")
			if err != nil {
				fmt.Println(err)
				continue
			}
			go func() {
				rd := bufio.NewReader(conn)
				fmt.Printf("server local addr  %v\n server remote addr %v \n", conn.LocalAddr(), conn.RemoteAddr())
				for {
					str, err := rd.ReadString('\n')
					if err != nil {
						fmt.Println(err)
						break
					}
					fmt.Println(str)
				}
			}()
		}
	}()
	//client
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	conn, _ := net.DialTCP("tcp", nil, addr)
	fmt.Printf("client local addr %v \n client remote addr %v \n ", conn.LocalAddr(), conn.RemoteAddr())
	conn2 := conn
	fmt.Printf("client local addr %v \n client remote addr %v \n ", conn2.LocalAddr(), conn2.RemoteAddr())

}
