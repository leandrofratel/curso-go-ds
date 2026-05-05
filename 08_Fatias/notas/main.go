package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var notas []float64
	notas := []float64{}

	for {

		var input_txt string
		fmt.Printf("Entre com sua nota: ")
		fmt.Scan(&input_txt)

		if input_txt == "" {
			break
		}

		nota, err := strconv.ParseFloat(input_txt, 64)
		if err != nil{
			fmt.Println("Entre com uma nota válida!")
		}

		notas = append(notas, nota)
	}

	fmt.Println(notas)

}