package main

import (
	"github.com/tatsushid/go-fastping"
	"net"
	"fmt"
	"os"
	"time"
)

const targetIP = "123.125.115.110"

func PingBaidu() {
	p:=fastping.NewPinger()
	ra,err:=net.ResolveIPAddr("ip4:icmp","163.com")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv= func(addr *net.IPAddr, duration time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT:%v\n",addr.String(),duration)
	}
	p.OnIdle= func() {
		fmt.Println("finsh")
	}
	p.RunLoop()
	err=p.Run()
	if err!=nil{
		fmt.Println(err)
	}
}
