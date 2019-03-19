package GORoutine

import (
	"fmt"
	"time"
)

func Run() {
	ch1 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			ch1 <- time.Now().String()
		}
	}()
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case <-time.After(time.Millisecond * 100):
			fmt.Print(".")
		}
	}
}
