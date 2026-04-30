package main

import (
	"fmt"
)

func main(){

	nome := "Leandro Fratel"

	for _, v := range nome {
		fmt.Println(string(v))
	}

}