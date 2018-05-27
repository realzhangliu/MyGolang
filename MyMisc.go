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

type t_face interface {
	Prt() string
	Check() bool
}

type t_face_struct struct {
}

func (s *t_face_struct) Prt() string {
	return fmt.Sprint("return a string")
}

func (s *t_face_struct) Check() bool {
	return true
}

//返回接口类型
func test_t_interface(t t_face) t_face {

	return &t_face_struct{}

}

