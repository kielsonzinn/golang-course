package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("hello world")
	escrever("programming in go")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
