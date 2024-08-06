package main

import (
	"fmt"
)

type Visitante struct {
	Nome  string
	Idade int
}

func main() {
	pessoas := []Visitante{
		Visitante{
			Nome:  "Clara",
			Idade: 27,
		},
		Visitante{
			Nome:  "Victor",
			Idade: 32,
		},
		Visitante{
			Nome:  "Dante",
			Idade: 1,
		},
	}

	for _, vis := range pessoas {
		// fmt.Println(i, vis)

		if vis.Idade >= 18 {
			HelloVisitante(vis.Nome)
		} else {
			fmt.Println("Você não tem idade para acessar esse evento...")
		}
	}
}

func HelloVisitante(name string) {
	fmt.Println("Hello, " + name)
}
