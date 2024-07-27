package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Não é bem herança da OO - #ficadica")

	p1 := pessoa{"João", "Pereira", 20, 178}
	e1 := estudante{p1, "Eng. Software", "USP"}

	fmt.Println(e1)

	fmt.Println(e1.nome)

}
