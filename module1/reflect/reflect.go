package main

import (
	"fmt"
	"reflect"
)

func main() {
	// basic type
	// myMap := make(map[string]string, 10)
	// myMap["a"] = "b"
	// t := reflect.TypeOf(myMap)
	// fmt.Println("type:", t)
	// v := reflect.ValueOf(myMap)
	// fmt.Println("value:", v)
	// struct
	myStruct := T{A: "a", B: "b"}
	v1 := reflect.ValueOf(myStruct)
	if v1.Kind().String() != "struct" {
		fmt.Print("Panic for v1 is not struct")
	}
	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v1.Field(i))
	}
	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}
	// 需要注意receive是struct还是指针
	result1 := v1.Method(0).Call(nil)
	fmt.Println("result:", result1)
	result2 := v1.Method(1).Call(nil)
	fmt.Println("result:", result2)
}

type T struct {
	A string
	B string
}

// 需要注意receive是struct还是指针
func (t T) String() string {
	return t.A + "a"
}

// 需要注意receive是struct还是指针
func (t T) Print() string {
	return t.B + "b"
}
