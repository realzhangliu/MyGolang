package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

const ratio =192

func main() {
	fmt.Println(os.Getwd())
	f, _ := os.Open("ImageOperation/T.jpg")
	defer f.Close()

	foutput, _ := os.Create("ImageOperation/output.jpg")
	defer foutput.Close()

	//str 是格式:jpeg
	im, str, err := image.Decode(f)
	fmt.Printf("Bounds is :%v\n", im.Bounds())
	fmt.Println("format is:" + str)
	if err != nil {
		log.Fatal(err)
	}
	im = resize.Resize(ratio, uint((ratio/im.Bounds().Max.X)*im.Bounds().Max.Y), im, resize.Lanczos3)
	fmt.Printf("output Bounds is :%v", im.Bounds())
	err = jpeg.Encode(foutput, im, nil)
	if err != nil {
		fmt.Println(err)
	}
}
