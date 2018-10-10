package Misc

import (
	"fmt"
	"github.com/gographics/gmagick"
	"os"
)

func RunProcess() {
	FuncOne()
}

func FuncOne() {
	fmt.Println(os.Getwd())
	im := gmagick.NewMagickWand()
	defer im.Destroy()
	err := im.ReadImage("too.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(im.GetImageColors())
	//im.WriteImage("too.tga")

}
