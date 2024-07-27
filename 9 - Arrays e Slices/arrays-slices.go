package main

import "fmt"

func main() {
	fmt.Println("Array e Slice")

	var array [5]int
	fmt.Println(array)

	array2 := [5]string{"Pos.1", "Pos.2", "Pos.3", "Pos.4", "Pos.5"}
	fmt.Println(array2)

	///slice
	slice := []int{2, 23, 4, 6}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3] //funciona como ponteiro
	fmt.Println(slice2)

	array2[1] = "Pos. 1 alterada"
	fmt.Println(slice2)

	fmt.Println("---------- Array interno")
	slice3 := make([]float32, 10, 11)

	slice3 = append(slice3, 5)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
