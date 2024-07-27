package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Mauricio",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "lagix",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "curso")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "peixes",
	}

	fmt.Println(usuario2)
}
