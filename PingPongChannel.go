package main

import (
	"fmt"
	"time"
)

type Ball struct {
	counter int
}

func RunPingPongGame() {
	table := make(chan *Ball)
	go Player("One", table)
	go Player("Two", table)
	table <- new(Ball)
	time.Sleep(time.Second * 1)
	<-table
	fmt.Println("game end")
}
func Player(s string, tables chan *Ball) {
	for {
		Ball := <-tables
		Ball.counter++
		fmt.Printf("%s hint times:%d\n", s, Ball.counter)
		time.Sleep(100 * time.Millisecond)

		tables <- Ball
	}
}

//channel of channel of channel

func testChannel() {

	go func() {
		go func() {
			go func() {

				fmt.Println("layer 3 start")
				time.Sleep(time.Second * 5)
				fmt.Println("layer 3 exit")

			}()
			fmt.Println("layer 2 start")
			time.Sleep(time.Second * 4)
			fmt.Println("layer 2 exit")
		}()
		fmt.Println("layer 1 start")
		time.Sleep(time.Second * 2)
		fmt.Println("layer 1 exit")
	}()
	fmt.Println("启动")
	fmt.Println("退出")
}
