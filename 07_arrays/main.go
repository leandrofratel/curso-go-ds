package main

import "fmt"

func main(){
	var y = [20]float64{1, 20, 42, 222}
	// fmt.Println(y)
	// fmt.Println("Primeiro elemento de Y:", y[0])
	// fmt.Println("Dois primeiro elementos de Y:", y[0:2])
	// fmt.Println("Terceiro em diante:", y[2:])

	// Mundando dados em um array
	y[18] = 1001
	y[19] = 2002

	// fmt.Println("Todos os elementos em Y:", y)
	// fmt.Println("Último elemento em Y:", y[19])
	// fmt.Println("Tamanho do array:", len(y))

	// Somar todos os elementos em Y
	var total float64
	for _, valor := range y{
		fmt.Println("Valor =", valor)
		total += valor
	}
	fmt.Println("Soma de Y=", total)
}