package Hash

import (
	"crypto/sha256"
	"fmt"
)

func NewHash() {
	Sha1Inst := sha256.New()
	Sha1Inst.Write([]byte("hello."))
	fmt.Printf("%v\n", Sha1Inst.Sum(nil))
	fmt.Printf("%x\n", Sha1Inst.Sum(nil))

}

func init() {
	fmt.Println("Initialization...")
}