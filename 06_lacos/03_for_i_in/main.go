package main

import (
	"fmt"
)

func main(){

	nome := "Leandro Fratel"

	for i := range nome {
		fmt.Println(string(nome[i]), i ,nome[i],)
	}

}