package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func Checkerr(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(0)
	}
}

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

//支持的RPC运算方法，所有方法以结构为基础
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port
	defer ln.Close()
	// run loop forever (or until ctrl-c)
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		//go handleConnection(conn)
		//rpc
		//go rpc.ServeConn(conn)
		go jsonrpc.ServeConn(conn)
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
