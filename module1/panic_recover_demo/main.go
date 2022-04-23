package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Main Func: Run start")
	defer func() {
		err := recover()
		if myErr, ok := err.(MyError); ok {
			fmt.Println("Main Func cache Error: ", myErr.Code, myErr.Messages)
		} else {
			fmt.Println("Main Func cache Error: ", err)
		}
	}()
	WriteFile("file.txt", "hello, Go")
}

type MyError struct {
	Code     uint
	Messages string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error Code: %d Error Messeages: %s", e.Code, e.Messages)
}

//
func WriteFile(fileName string, content string) {
	file, err := os.Open(fileName)
	defer func() {
		err := recover()
		switch err.(type) {
		case *os.PathError:
			// fmt.Println("WriteFile func cache Error: ", "file not exist.")
			panic(err)
		default:
			panic(MyError{Code: 1001, Messages: "文件不存在"})
		}
	}()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(content)
}
