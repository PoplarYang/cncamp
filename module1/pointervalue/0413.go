package main

import "fmt"

func main() {
	// var sum *int
	// sum = new(int) //分配空间
	// *sum = 98
	// fmt.Println(sum)

	//
	type Student struct {
		name string
		age  int
	}

	// 形式1
	var s *Student
	s = new(Student) //分配空间
	(*s).name = "dequan"
	// s.name = "dequan"
	fmt.Println(s)

	// 形式2
	s1 := &Student{"yang", 10}
	fmt.Println(s1)
}
