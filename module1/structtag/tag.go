package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Info struct {
	Name   string //name 是映射成mysql中的char类型还是varchar类型，还是text类型，即使能够说明，但是额外的信息max_length怎么表示
	Age    int    `json:"age,omitempty"` //如果有不想序列化到json解析的格式中的化价格"json：-"
	Gender string `json:"sex"`
}

//反射包

func main() {
	info := Info{
		Name:   "bobby",
		Gender: "nan",
	}
	re, _ := json.Marshal(info)
	fmt.Println(string(re))

	//通过反射包去识别type
	t := reflect.TypeOf(info)

	fmt.Println(t.Name())
	fmt.Println(t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) //获取结构体的每一个字段
		tag := field.Tag.Get("json")
		fmt.Printf("'%s'\n", tag)
	}

	//具体的引用，绝大多数情况下其实用不到反射，实际开发的项目中可能也会用到
}
