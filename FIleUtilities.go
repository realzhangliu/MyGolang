package main

import (
	"os"
	"fmt"
)

const (
	path     = "D:/MyGO/temp"
	temp_txt = "temp_txt.txt"
	temp_doc = "temp_doc.doc"
)

func check_panic(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string, debug_symbol bool) bool {
	exist := true
	if info, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	} else {
		fmt.Printf("File Name:%15s \nFile Size:%13.2fMB\n", info.Name(), float64(info.Size())/float64(1<<20))
	}
	return exist
}

func readFile(filename string) {
	fullFilename := path + "/" + temp_txt
	if checkFileIsExist(fullFilename, true) {
		f, err := os.Open(fullFilename)
		check_panic(err)
		defer f.Close()

		dat := make([]byte,5)

		for{
			if _,err:=f.Read(dat);err!=nil{
				break
			}else {
				fmt.Print(string(dat))

			}
		}


	}
}
