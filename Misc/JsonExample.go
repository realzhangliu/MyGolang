package Misc

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type response1 struct {
	Name string
	Age  int
	Time time.Time
}

func checkErro(e error) {
	if e!=nil{
		panic(e)
	}
}
func ExampleJson() {
	//The JSON package can automatically encode your custom data types.It will only include export fields in the encode output
	stru1 := response1{Name: "zl", Age: 999, Time: time.Now()}
	databytes, err := json.Marshal(stru1)
	checkErro(err)
	fi, err := os.OpenFile("data.json", os.O_CREATE|os.O_RDWR, 0666)
	defer fi.Close()
	fmt.Println(fi.Write(databytes))

	var dat map[string]interface{}

	json.Unmarshal(databytes, &dat)
	AgeFloat := dat["Age"].(float64)
	fmt.Println("Age:",AgeFloat)

	NameStr:=dat["Name"]
	fmt.Println(NameStr)


}
