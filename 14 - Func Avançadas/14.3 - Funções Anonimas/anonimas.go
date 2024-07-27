package main

import "fmt"

func main() {
	retorno := func(message string) string {
		return fmt.Sprintf("Msg Recebida: %s", message)
	}("Passando a mensagem")

	fmt.Printf(retorno)
}
