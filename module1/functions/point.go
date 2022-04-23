package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func PrintCat(cat Cat) {
	fmt.Println(cat)
}

func UpdateAge(cat *Cat) {
	cat.Age += 1
}

func main() {
	cat := Cat{Name: "yang", Age: 5}
	PrintCat(cat)
	UpdateAge(&cat)
	PrintCat(cat)
}
