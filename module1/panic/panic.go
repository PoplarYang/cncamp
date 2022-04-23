package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer func is called")
		if err := recover(); err != nil {
			fmt.Println("Recover error with error: ", err)
		}
	}()

	panic("a panic is trrigered")
	fmt.Println()
}
