package main

import "fmt"

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos.\n", nome, idade)
}

func main() {
	imprimirDados("Fernando", 29)
	fmt.Println(soma(3, 5))
}
