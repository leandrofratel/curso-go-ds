package main

import "fmt"

func main(){
	i := 1

	// Laço for sem condição de parada. Isso faz ele rodar continuamente.
	for {
		fmt.Println("Alguma coisa", i)
		i++
	}
}