package main

import (
	"cash-machine/calculate"
	"fmt"
)

func main() {
	var amount int
	fmt.Println("Digite o valor a ser sacado: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	notes, err := calculate.CalculateNotes(amount)
	first := true
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quantidade de notas: ")
		for denomination, count := range notes {
			if count > 0 {
				if !first {
					fmt.Printf(", ")
				}
			}
			if count == 1 {
				fmt.Printf("%d nota de %d\n", count, denomination)
			} else {
				fmt.Printf("%d notas de %d\n", count, denomination)
			}
		}
	}
}
