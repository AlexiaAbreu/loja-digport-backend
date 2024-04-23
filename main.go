package main

import "fmt"

func main() {
	//fmt.Printf("Esse é o catálogo: \n%+v\n", retornaCatalogo())
	nome := ""
	fmt.Scan("DIgite o nome do produto desejado: ", &nome)

	nomeDoProduto := retornaProdutoPeloNome(nome)
	fmt.Println(nomeDoProduto)

	StartServer()
}
