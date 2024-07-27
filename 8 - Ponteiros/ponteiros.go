package main

import "fmt"

func main() {
	fmt.Println("d")

	//PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel1 int
	var ponteiro *int
	fmt.Println(variavel1, ponteiro)

	variavel1 = 100
	ponteiro = &variavel1 //APONTANDO PARA O ENDEREÇO DE MEMORIA
	fmt.Println(variavel1, ponteiro)

	//DESREFERENCIAÇÃO
	fmt.Println(variavel1, *ponteiro)
}
