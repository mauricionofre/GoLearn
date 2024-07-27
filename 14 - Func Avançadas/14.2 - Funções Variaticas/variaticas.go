package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) string {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return texto + string(rune(total))
}

func main() {
	total := soma(1, 2, 3, 1, 2, 5, 4, 8, 7, 2, 1, 5)
	textoSomado := escrever("chupi", 1, 2, 3, 1, 2, 5, 4, 8, 7, 2, 1, 5)

	fmt.Println(total)
	fmt.Println(textoSomado)
}
