package main

import "fmt"

func main() {
	// a := []int{}
	// b := []int{1, 2, 3}
	// c := &a
	// a = append(b, 1)
	// fmt.Printf("%p %p\n", a, *c)

	a := make([]int, 10)
	// b := []int{1, 2, 3, 4}
	c := a
	// a = append(b)
	fmt.Printf("%p %p\n", a, c)
	// fmt.Printf("%d %d | %d %d\n", len(a), cap(a), len(c), cap(c))
}
