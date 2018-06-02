package main

import "fmt"

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
