package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGrupo sync.WaitGroup

	waitGrupo.Add(4)

	go func() {
		escrever("hello world")
		waitGrupo.Done()
	}()

	go func() {
		escrever("programming in go")
		waitGrupo.Done()
	}()

	go func() {
		escrever("Goroutine 3")
		waitGrupo.Done()
	}()

	go func() {
		escrever("Goroutine 4")
		waitGrupo.Done()
	}()

	waitGrupo.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
