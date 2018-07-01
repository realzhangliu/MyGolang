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
	"crypto/md5"
)

var projectPath string

func Start() {
	//fmt.Println(os.Args[0])
	//Path for assets
	var err error
	projectPath, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	projectPath = filepath.ToSlash(projectPath) + "/github.com/netldds/MyGolang/MyUploadServer"
	fmt.Println(projectPath)

	http.HandleFunc("/", syahelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("logout", logout)
	http.ListenAndServe(":9090", nil)
	fmt.Println("Server listing at: 127.0.0.1:9090")
	//http.ListenAndServe(":9090", http.FileServer(http.Dir(".")))

}
func logout(writer http.ResponseWriter, r *http.Request) {

}
func syahelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println(r.Form)
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Print("val:", value)
	}
	for _,v:=range r.Cookies(){
		fmt.Fprintf(w, "cookies:%v",v)
	}
	fmt.Fprintf(w,"%s","123")
	//http.Redirect(w,r,"http://www.baidu.com",http.StatusFound)
}

func uploadHandler(writer http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseMultipartForm(2 << 21)
		CheckHttpErrors(err, writer)
		m := r.MultipartForm
		files := m.File["uploadfile"]
		for _, value := range files {
			file, err := value.Open()
			CheckHttpErrors(err, writer)
			defer file.Close()
			md5str:=md5.New()
			io.WriteString(md5str,strconv.Itoa(time.Now().Second()))
			fullfilename:=fmt.Sprintf("%s/upload/%x%s",projectPath,md5str.Sum(nil),value.Filename)
			//dst, err2 := os.Create(projectPath + "/upload/" +  + value.Filename)
			dst, err2 := os.Create(fullfilename)
			defer dst.Close()
			CheckHttpErrors(err2, writer)
			_, err3 := io.Copy(dst, file)
			CheckHttpErrors(err3, writer)

			fmt.Fprintf(writer, "upload success.")

		}
	case "GET":
		t, err := template.ParseFiles(projectPath + "/uploadfiles.html")
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
		t, _ := template.ParseFiles(projectPath + "/login.html")
		log.Println(t.Execute(writer, nil))
	} else {
		r.ParseForm()
		log.Printf("username:%v", r.Form["username"])
		log.Printf("password:%v", r.Form["password"])
		log.Println(r.Form)
	}
}
