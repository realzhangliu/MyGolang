package Misc

import (
	"fmt"

	matrix "github.com/netldds/MyGolang/Matrix"
)

type MyError struct {
	id          int
	description string
}

func test() error {
	return &MyError{99, "qq"}
}

func (e *MyError) Error() string {
	e.id++
	return fmt.Sprintf("id:%d  description:%s", e.id, e.description)
}

type tFace interface {
	Prt() string
	Check() bool
}

type tFaceStruct struct {
}

func (s *tFaceStruct) Prt() string {
	return fmt.Sprint("return a string")
}

func (s *tFaceStruct) Check() bool {
	return true
}

//返回接口类型
func testTinterface(t tFace) tFace {

	return &tFaceStruct{}

}

func RunMatrixMutiply() {
	mt := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
		[]float32{7, 8, 9}}
	mv := [][]float32{
		{1, 2},
		{1, 2},
		{1, 2}}
	mt1, _ := matrix.New(mt)
	mt2, _ := matrix.New(mv)
	result, _ := mt1.Multiply(mt2)
	fmt.Println(*result)
}
