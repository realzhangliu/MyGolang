package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//GIN FrameWork APP
func main() {

	//DefaultAPI()
	//	Middleware()
	//renderData()
	ClockStream()
	//ServingData()
	//htmlRendering()
	//middleRun()
	//UsingBasicAuth()
	//Goroutinesmiddleware()
	//Encrypt()
	//AutoCert()
	//RunMultiServices()
	//GracefulShutdown()
	//BindFormData()
	//ServerPush()
	//queryDB()
	//fmt.Println(len("9f1431e28609f22fb5e6fcd9f713e8d6"))
	//WebSocketImplementation()
	//StartSrv()
	//input()
	//ForwardProxyStart()
}
func input() {
	rd := bufio.NewReader(os.Stdin)
	for {
		delim, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(delim)
	}
}
