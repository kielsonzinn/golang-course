package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	type usuario struct {
		nome  string
		idade int
	}

	usuario1 := usuario{"teste", 2}

	generica(1)
	generica("teste")
	generica(2.24)
	generica(false)
	generica(usuario1)

	mapa := map[interface{}]interface{}{
		int(1):       "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)

}
