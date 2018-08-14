package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type handle struct {
	reverseProxy string
}

func (e *handle) ServeHTTP(rep http.ResponseWriter, req *http.Request) {
	remote, _ := url.Parse(e.reverseProxy)
	fmt.Printf("Request host is:%s\n Remote host is:%s\n", req.Host, remote.Host)

	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}
		// http://127.0.0.1/s?wd=b
		return dialer.DialContext(ctx, network, "220.181.57.216:80")
		//return dialer.DialContext(ctx, network, addr)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	req.Host = remote.Host

	proxy.ServeHTTP(rep, req)
}
func StartSrv() {
	h := &handle{
		reverseProxy: "http://www.baidu.com",
	}
	srv := http.Server{
		Handler: h,
		Addr:    ":80",
	}
	srv.ListenAndServe()
}
