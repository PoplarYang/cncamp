package main

import "fmt"

func delete(s []int, index int) []int {
	if index >= len(s) {
		return nil
	}
	return append(s[0:index], s[index+1:]...)

}

func main() {
	for i := 1; i <= 3; i++ {
		v := 1
		fmt.Println(&v)
	}
}
