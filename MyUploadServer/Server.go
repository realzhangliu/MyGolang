package MyUploadServer

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"io"

	"strconv"
	"time"

	"html/template"
	"path/filepath"
)

var projectPath string

func Start() {
	//fmt.Println(os.Args[0])
	//Path for assets
	var err error
	projectPath,err=os.Getwd()
	if err!=nil{
		log.Println(err)
	}
	projectPath=filepath.ToSlash(projectPath)+"/src/MyGolang/MyUploadServer"
	fmt.Println(projectPath)

	http.HandleFunc("/", syahelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":9090", nil)
	//http.ListenAndServe(":9090", http.FileServer(http.Dir(".")))


}
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

func uploadHandler(writer http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseMultipartForm(10000)
		CheckHttpErrors(err, writer)
		m := r.MultipartForm
		files := m.File["uploadfile"]
		for _, value := range files {
			file, err := value.Open()
			CheckHttpErrors(err, writer)
			defer file.Close()
			dst, err2 := os.Create(projectPath+"/upload/" + strconv.Itoa(time.Now().Second()) + value.Filename)
			defer dst.Close()
			CheckHttpErrors(err2, writer)
			_, err3 := io.Copy(dst, file)
			CheckHttpErrors(err3, writer)
			for i := 0; i < 3; i++ {
				fmt.Fprintf(writer, "processing %d ...%s", i, value.Filename)
			}
		}
	case "GET":
		t, err := template.ParseFiles(projectPath+"/uploadfiles.html")
		CheckHttpErrors(err, writer)
		t.Execute(writer, nil)

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func CheckHttpErrors(e error, writer http.ResponseWriter) {
	if e != nil {
		http.Error(writer, e.Error(), http.StatusInternalServerError)
		//os.Exit(1)
	}
}
func login(writer http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles(projectPath+"/login.html")
		log.Println(t.Execute(writer, nil))
	} else {
		r.ParseForm()
		log.Printf("username:%v", r.Form["username"])
		log.Printf("password:%v", r.Form["password"])
		log.Println(r.Form)
	}
}
