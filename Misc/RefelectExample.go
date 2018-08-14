package main

import (
	"reflect"
	"fmt"
)

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

//insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
func createQuery(q interface{})  {

	q=employee{name:"zl",
	id:5,
	address:"tianhong",
	salary:18000,
	country:"CHN"}
	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v:=reflect.ValueOf(q)
		query:=fmt.Sprintf("Insert into %s values(",reflect.TypeOf(q).Name())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query=fmt.Sprintf("%s%d",query,v.Field(i).Int())
				}else{
					query=fmt.Sprintf("%s,%d",query,v.Field(i).Int())
				}
			case reflect.String:
				if i==0{
					query=fmt.Sprintf("%s\"%s\"",query,v.Field(i).String())
				}else {
					query=fmt.Sprintf("%s,\"%s\"",query,v.Field(i).String())
				}
			}
		}
		query=fmt.Sprintf("%s)",query)
		fmt.Println(query)
	}
}
