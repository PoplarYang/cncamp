package main

import (
	"log"
	"testing"
)

func TestDelete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	index := 4
	// result := []int{1, 2, 3, 4}
	slice = delete(slice, index)
	log.Println(slice)
}
