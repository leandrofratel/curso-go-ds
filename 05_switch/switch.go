package main

import "fmt"

func main(){
	fruta := "morango"

	switch fruta {

	case "amora":
		fmt.Println("Você escolheu amora! adoro essa fruta!")

	case "maça":
		fmt.Println("Eu já não gosto tanto de maçãs.")

	case "morango":
		fmt.Println("Gosto ainda mais de moragos!")

	default:
		fmt.Println("Informe o nome de uma fruta")
	}
}