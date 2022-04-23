package main

import "fmt"

func increase(x, y int) {
	fmt.Println("Increase result is: ", x+y)
}

func decrease(x, y int) {
	fmt.Println("Decrease result is: ", x-y)
}

func DoOperation(original int, f func(int, int)) {
	f(original, 1)
}

func main() {
	DoOperation(2, increase)
	DoOperation(2, decrease)
}
