package main // todo arquivo por padrão terá um package

import (
	"fmt"
	"reflect"
)

func main() {
	CriarVariaveis()
	fmt.Print("Imprimindo o return da funcao Soma \n", Soma(10, 10), "\n")
	fmt.Print("Imprimindo o return da funcao Hello \n", Hello(), "\n")

}

func Soma(a, b int) int {
	return a + b
}

func CriarVariaveis() {
	fmt.Printf("Esta eh a funcao CriarVariaveis \n")
	inteiro := 1
	texto := "bla bla bla"
	flutuante := 3.5
	booleano := true

	array := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	array_strings := []string{}
	fmt.Printf("Isto eh um print de um array \n")

	for _, item := range array {
		fmt.Println(item)
	}

	fmt.Printf("Isto eh um print dos tipos de variaveis \n")

	fmt.Printf("inteiro: %s\n", reflect.TypeOf(inteiro))
	fmt.Printf("texto: %s\n", reflect.TypeOf(texto))
	fmt.Printf("booleano: %s\n", reflect.TypeOf(flutuante))
	fmt.Printf("booleano: %s\n", reflect.TypeOf(booleano))
	fmt.Printf("array de inteiros: %s\n", reflect.TypeOf(array))
	fmt.Printf("array de strings: %s\n", reflect.TypeOf(array_strings))

}

func Hello() string {
	return ("Hello World")
}
