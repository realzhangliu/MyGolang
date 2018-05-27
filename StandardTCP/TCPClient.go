package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	CheckErr(err)
	tcpconn, err2 := net.DialTCP("tcp", nil, tcpaddr)
	CheckErr(err2)
	defer tcpconn.Close()

	go handleMessageFromServer(tcpconn)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		//tcpconn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := tcpconn.Write([]byte(text))
		if text=="q\n"{
			os.Exit(0)
		}
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}

}
func handleMessageFromServer(conn *net.TCPConn) {
	reply := make([]byte, 1024)
	for {
		_, err := conn.Read(reply)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(reply)
	}
}
