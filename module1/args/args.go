package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "input name")
	flag.Parse()
	fmt.Printf("%s", *name)
}
