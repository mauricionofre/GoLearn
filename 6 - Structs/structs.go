package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	u.idade = 30
	u.nome = "Mauricio"
	fmt.Println(u)

	usuario1 := usuario{"Onofre", 35}

	fmt.Println(usuario1)

	usuario2 := usuario{nome: "Pereira"}

	fmt.Println(usuario2)

}
