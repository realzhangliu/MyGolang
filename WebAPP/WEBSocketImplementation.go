package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/websocket"
)

var DBserver string = "ml:0000@tcp(127.0.0.1)/person"

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func WebSocketImplementation() {
	r := gin.Default()
	//db, err := sql.Open("mysql", DBserver)
	r.GET("/ws", func(c *gin.Context) {
		handler := websocket.Handler(EchoServer)
		handler.ServeHTTP(c.Writer, c.Request)
	})
	//checkErr(err)
	r.Run(":80")
}

func EchoServer(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, "123"); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
