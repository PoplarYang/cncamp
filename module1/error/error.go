package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建error的两种方式
	notFoundError := fmt.Errorf("this is a user defind error.")
	fmt.Println(notFoundError)
	permissionDeniedError := errors.New("this is a permission denied error.")
	fmt.Println(permissionDeniedError)
}
