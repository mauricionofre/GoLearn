package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	fmt.Println(numero)

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Maior que zero")
	}

}
