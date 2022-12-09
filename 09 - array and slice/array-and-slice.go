package main

import (
	"fmt"
)

func main() {

	fmt.Println("arrays and slice")

	var array1 [5]string
	array1[0] = "pos 01"
	fmt.Println(array1)

	array2 := [5]string{"pos 01", "pos 02", "pos 03", "pos 04", "pos 05"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "pos changed"
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("----------------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
