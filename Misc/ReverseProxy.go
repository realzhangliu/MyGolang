package Misc

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func ReverseProxyStart() {
	http.HandleFunc("/new", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "new function")
	})

	u1, _ := url.Parse("http://www.baidu.com")
	http.Handle("www.baidu.com", httputil.NewSingleHostReverseProxy(u1))
	http.ListenAndServe(":80", nil)
}


