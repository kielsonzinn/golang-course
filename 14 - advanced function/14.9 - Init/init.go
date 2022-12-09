package main

import "fmt"

var n int

func main() {
	fmt.Println("func main")
	fmt.Println(n)
}

func init() {
	fmt.Println("func init")
	n = 10
}
