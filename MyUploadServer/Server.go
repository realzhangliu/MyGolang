package MyUploadServer

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/template"
)

func syahelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Print("val:", value)
	}
	fmt.Fprintf(w, "Hello wangwang!")
}
func Start() {
	fmt.Println(os.Args[0])
	http.HandleFunc("/", syahelloName)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
func login(writer http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./MyGolang/MyUploadServer/login.html")
		log.Println(t.Execute(writer, nil))
	} else {
		r.ParseForm()
		log.Printf("username:%v", r.Form["username"])
		log.Printf("password:%v", r.Form["password"])
		log.Println(r.Form)
	}
}
