package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

const FILE_DIR = "DATA"

const URL = "http://localhost:8090/api/v1/s/netio/files"

type Gatlin struct {
	locker sync.Mutex
}
type Bullets struct {
	Magazine map[string]os.FileInfo
}

func (b *Bullets) Shot(name string) {
	clinet := http.Client{}

	req, _ := http.NewRequest("POST", URL, nil)
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDU5ODgyNjcsImlkIjoiNjNhNDBiNWYzMjBmNDRjY2I2NjM5YTQ0YjBjZTgyMzIiLCJpc3MiOiJkYXhpYW5neXVuLmNvbSIsIm9yaWdfaWF0IjoxNTQ1MTI0MjY3fQ.ZdE80JmyWbZh1kQZ54ONsZuWyIxLFaQ1slwAEyaVOaM")
	req.Header.Set("Upload-Length", strconv.Itoa(int(b.Magazine[name].Size())))
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Tus-Resumable", "1.0.0")
	req.Header.Set("Upload-Metadata", "filename"+" "+base64.StdEncoding.EncodeToString([]byte(name)))
	resp, err := clinet.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println(resp.Status)
	//fmt.Println(resp.Header)
	if resp.Status != "201 Created" {
		return
	}
	location := resp.Header["Location"][0]
	u, _ := url.Parse(location)
	paths := strings.Split(u.Path, "/")

	fileID := paths[len(paths)-1]

	//fmt.Println(fileID)

	var ContentLength int64 = 1024 * 100

	var UploadOffset int64 = 0

	fdata, _ := os.Open(path.Join(FILE_DIR, name))
	defer fdata.Close()
	n := 0
	for true {
		specifiedData := make([]byte, ContentLength)
		n, _ = fdata.Read(specifiedData)
		if n == 0 {
			fmt.Printf("%s is done \n", fileID)
			break
		}
		req, _ := http.NewRequest("PATCH", URL+"/"+fileID, bytes.NewBuffer(specifiedData[:n]))
		req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDU5ODgyNjcsImlkIjoiNjNhNDBiNWYzMjBmNDRjY2I2NjM5YTQ0YjBjZTgyMzIiLCJpc3MiOiJkYXhpYW5neXVuLmNvbSIsIm9yaWdfaWF0IjoxNTQ1MTI0MjY3fQ.ZdE80JmyWbZh1kQZ54ONsZuWyIxLFaQ1slwAEyaVOaM")
		req.Header.Set("Content-Type", " application/offset+octet-stream")
		req.Header.Set("Upload-Offset", strconv.FormatInt(UploadOffset, 10))
		req.Header.Set("Tus-Resumable", " 1.0.0")
		req.ContentLength = int64(n)

		resp, err := clinet.Do(req)

		//defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Println(resp.Status)
		if resp.Status == "204 No Content" {
			UploadOffset, _ = strconv.ParseInt(resp.Header.Get("Upload-Offset"), 10, 64)
		}
	}

}
func main() {
	fmt.Println(os.Getwd())
	bullets := Initialization_FILES()

	//bullets.Shot("globe.png")

	for _, v := range bullets.Magazine {
		go bullets.Shot(v.Name())
	}
	select {}
}

func Initialization_FILES() *Bullets {
	var bullets Bullets
	bullets.Magazine = make(map[string]os.FileInfo)
	Finfos, err := ioutil.ReadDir(FILE_DIR)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(Finfos); i++ {
		if Finfos[i].IsDir() {
			continue
		}
		bullets.Magazine[Finfos[i].Name()] = Finfos[i]
	}
	return &bullets
}
