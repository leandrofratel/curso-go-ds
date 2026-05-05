package main

import "fmt"

func main() {

	// ARRAY
	var x [100]int

	fatia := x[:10]
	fmt.Printf("X = %T | fatia = %T\n", x, fatia)

	// tamanho das fatias
	fmt.Println("Tamanho da fatia:",len(fatia))

	y:=[]int{}

	for i:=1; i<=200; i++{
		y = append(y, i)
		fmt.Println(len(y))
	}
	fmt.Println(y)
}