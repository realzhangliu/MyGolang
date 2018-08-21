package Misc

import (
	"net/http"
	"fmt"
	"net/url"
	"net/http/httputil"
)

func ReverseProxyStart(){
	http.HandleFunc("/new", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"new function")
	})

	u1,_:=url.Parse("http://www.baidu.com")
	http.Handle("www.baidu.com",httputil.NewSingleHostReverseProxy(u1))
	http.ListenAndServe(":80",nil)
}
