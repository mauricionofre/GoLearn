package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo")
	escrever("Segunda parada")

}
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
