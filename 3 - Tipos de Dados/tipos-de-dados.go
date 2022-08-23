package main

import (
	"errors"
	"fmt"
)

func main() {

	// Numeros Inteiros
	var numero int64 = 10000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Numeros Reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Char
	char := 'A'
	fmt.Println(char)

	var texto := 5
	fmt.Println(texto)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
