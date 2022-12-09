package main

import "fmt"

func main() {

	var var1 string = "var 1"
	fmt.Println(var1)

	var2 := "var 2"
	fmt.Println(var2)

	var (
		var3 string = "var 3"
		var4 string = "var 4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var 5", "var 6"

	fmt.Println(var5, var6)

	const const1 string = "const 1"
	fmt.Println(const1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
