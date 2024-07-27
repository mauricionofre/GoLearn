package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 = numero1 + numero2
	fmt.Println(soma2)

	var variavel1 string = "string"
	variavel2 := "string2"

	fmt.Println(variavel1, variavel2)

	//não existe operador ternario
	//texto := numero1 > 5 ? "Maior que 5": "menor que 5"

}
