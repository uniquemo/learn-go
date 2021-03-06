package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func getTagInfo() {
	p := reflect.TypeOf(Person{Name: "person1", Age: 18})

	name, _ := p.FieldByName("Name")
	fmt.Printf("%q\n", name.Tag)

	age, _ := p.FieldByName("Age")
	fmt.Printf("%q\n", age.Tag)
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

	getTagInfo()
}
