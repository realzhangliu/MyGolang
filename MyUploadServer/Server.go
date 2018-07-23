package MyUploadServer

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"io"

	"strconv"
	"time"

	"crypto/md5"
	"encoding/json"
	"errors"
	"html/template"
	"net/rpc"
	"net/url"
	"path/filepath"
)

var projectPath string

type usertype struct {
	ID  string `json:id`
	Pwd string `json:"password"`
}
type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}
//支持的RPC运算方法，所有方法以结构为基础
type Arith int

//支持多个不同结构方法
type Brith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
func (e *Brith) Add(in *int, out *int) error {
	*out = *in + 1
	return nil
}

func Start() {
	//fmt.Println(os.Args[0])
	//PRC
	arith := new(Arith)
	brith := new(Brith)
	rpc.Register(brith)
	rpc.Register(arith)
	rpc.HandleHTTP()

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
	http.ListenAndServe(":80", nil)
	fmt.Println("Server listing at: 127.0.0.1:80")
	//http.ListenAndServe(":9090", http.FileServer(http.Dir(".")))

}
func logout(writer http.ResponseWriter, r *http.Request) {

}
func syahelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header.Get("Content-Type"))
	//fmt.Println(r.Form)
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key:", key)
		fmt.Println("val:", value)
	}
	for _, v := range r.Cookies() {
		fmt.Fprintf(w, "cookies:%v\n", v)
		fmt.Println(v)
	}
	//MaxAge是秒单位
	//Name是服务器指定的固定名字，Value是分每个客户端分配的唯一标识
	cookie := http.Cookie{Name: "MyCookieName", Value: url.QueryEscape("ValueAsSID"), Path: "/", HttpOnly: true, MaxAge: 3600}
	//发送COOKIE给客户端
	http.SetCookie(w, &cookie)
	var out usertype
	json.NewDecoder(r.Body).Decode(&out)
	fmt.Println(out)
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
			md5str := md5.New()
			io.WriteString(md5str, strconv.Itoa(time.Now().Second()))
			fullfilename := fmt.Sprintf("%s/upload/%x%s", projectPath, md5str.Sum(nil), value.Filename)
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
