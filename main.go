package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
}

func IsEven(i int) bool {
	fmt.Println(i % 2)
	if i%2 == 0 {
		return true
	}
	return false
}
