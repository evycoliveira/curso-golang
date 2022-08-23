package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// n1 e n2 sao do tipo int8
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	var fun = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f()
	resultado := fun("Texto da função fun.")
	fmt.Println(resultado)

	// Duas variaveis que recebem uma funcao que retorna dois valores
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// ResultadoSubtracao e ignorado, mesmo seu retorno sendo feito
	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}
