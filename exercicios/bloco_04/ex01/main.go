// Faça um programa que conte quantas vezes a letra 'a' aparece em uma palavra

package main

import (
	"fmt"
	"strings"
)

func main(){

	var palavra string
	fmt.Println("Entre com uma palavra: ")
	fmt.Scanf("%s", &palavra)

	// Garanti que tudo estará no minusculo
	palavra = strings.ToLower(palavra)

	contador := 0
	for _, letra := range palavra{
		if string(letra) == "a" {
			contador ++
		}
	}

	fmt.Printf("A palavra '%s' tem '%d' As", palavra, contador)
}