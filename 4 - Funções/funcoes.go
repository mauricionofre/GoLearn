package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (som int8, sub int8) {
	soma := somar(n1, n2)
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 10)
	fmt.Println(soma)

	var fu = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	var resultado = fu("jabacule")
	fmt.Println(resultado)

	resultadoSoma, resultadoSub := calculosMatematicos(10, 2)

	fmt.Println(resultadoSoma, resultadoSub)
}
