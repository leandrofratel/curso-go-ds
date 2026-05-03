/*
Faça um programa que receba uma quantidade indefinida e valores correspondentes a 'saldo em conta',
mas quando o usuário aperta "enter" sem digitar valor algum, o programa para de receber valores,
e exibe a soma de todos os valores digitados anteriormente.
*/
package main

import (
	"bufio"		// -> Fornece um buffer para leitura dos dados.
	"fmt"		// -> Formata as mensagens do terminal.
	"os"		// -> Acessa entradas do teclado.
	"strconv" 	// -> Converte de texto para número.
	"strings"	// -> Remove espaços e quebra ed linha.
)

func main() {

	// variavel que recebe um comando do teclado
	enter := bufio.NewReader(os.Stdin)
	var total float64

	fmt.Printf("Digite os saldos em conta (Enter vazio para encerrar): ")

	for {
		fmt.Printf("Saldo: ")
		entrada, _ := enter.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		// parada do laço
		if entrada == ""{
			break
		}

		// converte string para float64
		valor, erro := strconv.ParseFloat(entrada, 64)
		if erro != nil { // Verifica se há um erro
			fmt.Println("Valor inválido, Digite um número.")
			continue
		}

		total += valor
	}

	fmt.Printf("A soma de todos os saldos é: %.2f\n", total)

}