package main

import "fmt"

func main() {
	myArray := [5]int{0, 1, 2, 3, 4}
	mySlice := myArray[1:3]
	fmt.Printf("my slice %v", mySlice)
}

func DeleteByIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
