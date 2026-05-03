// Faça um programa que receba 4 alturas e usando um laço, realiza a soma delas.

package main

import "fmt"

func main() {

	var altura, total float64

	for i := 1; i <= 4; i++ {
		fmt.Printf("Entre com %dª altura: ", i)
		fmt.Scan(&altura)
		total += altura
	}

	fmt.Printf("A soma de todas as alturas é: %.2f\n", total)

}