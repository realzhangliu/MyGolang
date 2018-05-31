package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Checkerr(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(0)
	}
}
func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	defer ln.Close()
	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("remote client from" + remoteAddr)

	reader := bufio.NewReader(conn)
	for {
		data, err3 := reader.ReadString('\n')
		if err3 != nil {
			break
		}
		log.Printf("%v -> %s", conn.RemoteAddr(), data)
	}
	//for {
	//	message := make([]byte, 1024)
	//	i, err2 := conn.Read(message)
	//	if err2 != nil {
	//		break
	//	}
	//	log.Printf("%v -> %s",i,string(message))
	//}
	//scanner := bufio.NewScanner(conn)
	//for {
	//	ok := scanner.Scan()
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(scanner.Text())
	//}

}
