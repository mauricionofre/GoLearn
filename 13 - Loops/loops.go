package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i: ")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "lucas"}

	for i, value := range nomes {
		fmt.Println(i, value)
	}

	for i, value := range "PALAVRA" {
		fmt.Println(i, string(value))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for index, value := range usuario {
		fmt.Println(index, value)
	}
}
