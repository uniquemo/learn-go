package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

// struct tag 的应用：
func Print(obj interface{}) error {
	v := reflect.ValueOf(obj) // 取 Value

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析 label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}

func main() {
	person := Person{
		Name: "MING",
		Age:  18,
	}

	Print(person)
}
